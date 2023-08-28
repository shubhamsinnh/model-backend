import http from "k6/http";
import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';
import {
  check,
  group,
  sleep
} from "k6";

import {
  FormData
} from "https://jslib.k6.io/formdata/0.0.2/index.js";
import {
  randomString
} from "https://jslib.k6.io/k6-utils/1.1.0/index.js";

import {
  genHeader, genHeaderwithJwtSub,
} from "./helpers.js";

import * as constant from "./const.js"

const model_def_name = "model-definitions/local"

export function GetModelCard() {
  // Model Backend API: Get model card
  let resp = http.request("GET", `${constant.mgmtApiPrivateHost}/v1alpha/admin/users/${constant.defaultUserId}`, {}, {
    headers: genHeader(`application/json`),
  })
  let userUid = resp.json().user.uid

  let fd_cls = new FormData();
  let model_id = randomString(10)
  let model_description = randomString(20)
  fd_cls.append("id", model_id);
  fd_cls.append("description", model_description);
  fd_cls.append("model_definition", model_def_name);
  fd_cls.append("content", http.file(constant.cls_model, "dummy-cls-model.zip"));
  {
    let createClsModelRes = http.request("POST", `${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/multipart`, fd_cls.body(), {
      headers: genHeaderwithJwtSub(`multipart/form-data; boundary=${fd_cls.boundary}`, userUid),
    })

    // Check model creation finished
    let currentTime = new Date().getTime();
    let timeoutTime = new Date().getTime() + 120000;
    while (timeoutTime > currentTime) {
      let res = http.get(`${constant.apiPublicHost}/v1alpha/${createClsModelRes.json().operation.name}`, {
        headers: genHeader(`application/json`),
      })
      if (res.json().operation.done === true) {
        break
      }
      sleep(1)
      currentTime = new Date().getTime();
    }

    group(`Model Backend API: Get model card [with "jwt-sub" header]`, function () {
      check(http.get(`${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${model_id}/readme`, {
        headers: genHeaderwithJwtSub(`application/json`, uuidv4()),
      }), {
        [`[with random "jwt-sub" header] GET /v1alpha/models/${model_id}/readme response status 401`]: (r) =>
          r.status === 401,
      });

      check(http.get(`${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${model_id}/readme`, {
        headers: genHeaderwithJwtSub(`application/json`, userUid),
      }), {
        [`[with default "jwt-sub" header] GET /v1alpha/models/${model_id}/readme response status 200`]: (r) =>
          r.status === 200,
      });

      currentTime = new Date().getTime();
      timeoutTime = new Date().getTime() + 120000;
      while (timeoutTime > currentTime) {
        let res = http.get(`${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${model_id}/watch`, {
          headers: genHeader(`application/json`),
        })
        if (res.json().state !== "STATE_UNSPECIFIED") {
          break
        }
        sleep(1)
        currentTime = new Date().getTime();
      }

      // clean up
      check(http.request("DELETE", `${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${model_id}`, null, {
        headers: genHeaderwithJwtSub(`application/json`, userUid),
      }), {
        [`[with default "jwt-sub" header] DELETE clean up response status 204`]: (r) =>
          r.status === 204
      });
    });
  }
}
