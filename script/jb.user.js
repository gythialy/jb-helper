// ==UserScript==
// @name         jb-helper
// @namespace    https://github.com/gythialy/jb-helper
// @version      1.0
// @license MIT
// @description  Add a button on the plugin homepage and click to get the plugin activation code
// @author       novice.li
// @match        https://plugins.jetbrains.com/*
// @grant        GM_setClipboard
// @grant        GM_addStyle
// @grant        window.onurlchange
// @require https://cdnjs.cloudflare.com/ajax/libs/jsrsasign/10.5.27/jsrsasign-all-min.js
// @downloadURL https://https://raw.githubusercontent.com/gythialy/jb-helper/refs/heads/main/script/jb.user.js
// @updateURL https://https://raw.githubusercontent.com/gythialy/jb-helper/refs/heads/main/script/jb.user.js
// ==/UserScript==

const pemEncodedKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`;

const pemEncodedCrt = `-----BEGIN CERTIFICATE-----
MIIEuDCCAqCgAwIBAgIQehDN5nmBiuq8RGGmzOHc/zANBgkqhkiG9w0BAQsFADAY
MRYwFAYDVQQDEw1KZXRQcm9maWxlIENBMB4XDTI0MTAwNDAzMjAzM1oXDTM0MTAw
MzAzMjAzM1owGDEWMBQGA1UEAxMNSmV0UHJvZmlsZSBDQTCCAiIwDQYJKoZIhvcN
AQEBBQADggIPADCCAgoCggIBAOF8P6DPrWLldByoybs6rY3j09FMloRICY17aIEC
LYAmqPDUUAi7I9uxwMunFMXwcQVEUng3odh2FgZDAFpbjMzgdPIaKT97S78mnf/n
TI0wysHH04vRzeehfW9XX7vtjca1PmMYpuLhvW2al51xODxx5zipzNtAneLsaejC
zazR0JJLKoWe7FV2NWdFGzxRD0AUWSMjqn94uE6iu3Ci98rTxeHabVT7TzswWPNK
UsKiVzVbHbneoTl4UY5RQWJp7UfbRSCQKHbFGuTI7dDPJyLNruevLKDbK/LW+yem
2JeTb/2c6JRgQZZZK7IqRhaG7zK3TQsjmQmum5a7WD539lx78xKm1pKjLimveOwX
DFfaWSHWy5zR2MEXQxCNwbRPB2WRdt2hoqgcaKn5LeizDSaCwPJYOKXJFwTonn5R
XqhVmsMKkxZPss14+2VzbP/B2M/N4i6pyd+uYWHYF+myX9yfPbkcCA5tTfAwEr2L
H5iFlsv3WWS9sAkHl8vpndiNOXFzVPB0iPQSQghgFpk85OMhK0DnOsS2zrp+IBPu
Nw586V0GZrWuWy8jrIStur38UMrz2wAWwEXcRCIRkmUlHgRyuJn4CUT0CqHcAegN
/Mrpsf4XQRRPxb8w4k7yL5JpPEtajh+Sb5a7i6xynljIiO4+r4Ax84R6a4LQxgS3
vVv5AgMBAAEwDQYJKoZIhvcNAQELBQADggIBAEe2q6YsTPseWqU8av7d6VZH+Rff
8l5vYQSPgdAGLNVGW7R2V8t+1drp7NQepFaRkGkivUcjNKqa5sDwWUlPXCcboVdj
su1+SXxvb0G6VqI3+bCmd7oJtiewpOOxVek/VzNifKsJurbA9OUFrqkbgkqshxSW
i/nAVEsgft5dRu+dziM8847Zdc4omG8er0ei3gWK1ag7YUeodToba1VtrUS+aUtF
IZrnPFW1wlhyoUAtxKa57v/4kzhqSggbrVSvkprvoIJxVYgjoIt28Q3Nx09qmY68
cQP7ebhL3a4GltUd5isls1FHEJJaZUiroI2XQ7Wc+a/SrdeoirdAjadWJzF9A0PU
5+3DmGoEwQZliwWJg+/xCubROtUowWJqUXU9tqLSjOwNVJzpmTxv1asrvgYm/dTH
kuxU7OTBV80Xd6OJFpjOvWRsvE5XBvIFytCxOsQUZDlaW7arZkkvYOw0cN0vMTlU
HRiBcheCq76JPNx/u897TKRIXYeTQEyCcAKKIvi+YGpz15ALITNAUpJP+bz09r//
XflLX/bDwoKa6nq6k02xV6k1N3iDUyfLOnGefzplvQU9cW4wswGdhZFc0eUVSGtd
tcD9+XWm/y6wo0faHGE5gBKPvlhe+KyGr+GeDtAshtBiBd79/HVODfplxnLUdY9T
f+fBKDxCW0yk05Mo
-----END CERTIFICATE-----`;

function injectStyles() {
  GM_addStyle(`
        .jetbra-button {
            background-color: #04AA6D; border: none; color: white; padding: 8px 24px;
            text-align: center; text-decoration: none; display: inline-block;
            border-radius: 16px; box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2), 0 6px 20px 0 rgba(0,0,0,0.19);
            transition-duration: 0.4s;
            margin-left:5px
        }
        .jetbra-button:hover { background-color: #057e47; color: white; }
    `);
}

async function findElementWithRetry(cssSelector) {
  const maxAttempts = 50;
  for (let attempts = 0; attempts < maxAttempts; attempts++) {
    const element = document.querySelector(cssSelector);
    if (element) {
      return element;
    }
    await new Promise((resolve) => setTimeout(resolve, 100));
  }
  throw new Error(
    `Element with selector '${cssSelector}' not found after ${maxAttempts} attempts.`,
  );
}

function pem2base64(pem) {
  return pem
    .replace(/-----BEGIN.*?-----/, "")
    .replace(/-----END.*?-----/, "")
    .replace(/[\r\n]+/g, "");
}

function arrayBufferToBase64(buffer) {
  return btoa(
    [...new Uint8Array(buffer)].map((b) => String.fromCharCode(b)).join(""),
  );
}

function base64ToArrayBuffer(base64) {
  return Uint8Array.from(atob(base64), (c) => c.charCodeAt(0)).buffer;
}

function importPrivateKeyFromPem(pemKey) {
  const key = KEYUTIL.getKey(pemKey);
  return key;
}

function signData(privateKey, data) {
  const sig = new KJUR.crypto.Signature({ alg: "SHA1withRSA" });
  sig.init(privateKey);
  sig.updateString(data);
  const sigValueHex = sig.sign();
  return hextob64(sigValueHex);
}

function genLicenseId() {
  const CHARSET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  return Array.from({ length: 10 }, () => {
    let idx = Math.floor(Math.random() * CHARSET.length);
    return CHARSET[idx];
  }).join("");
}

function buildLicensePartJson(productCode, licenseId, twoYears) {
  let formattedDateTwoYears = "2099-08-01";
  if (twoYears) {
    let futureDateTwoYears = new Date();
    futureDateTwoYears.setFullYear(futureDateTwoYears.getFullYear() + 1);
    formattedDateTwoYears = futureDateTwoYears.toISOString().split("T")[0];
  }

  let fallbackDate = formattedDateTwoYears;
  let paidUpTo = formattedDateTwoYears;

  return JSON.stringify({
    licenseId: licenseId,
    licenseeName: "Mercury Labs",
    assigneeName: "Goren G",
    assigneeEmail: "",
    licenseRestriction: "",
    checkConcurrentUse: false,
    products: [
      {
        code: productCode,
        fallbackDate: fallbackDate,
        paidUpTo: paidUpTo,
        extended: false,
      },
    ],
    metadata: "0120230102PPAA013009",
    hash: "41472961/0:1563609451",
    gracePeriodDays: 7,
    autoProlongated: true,
    isAutoProlongated: true,
  });
}

async function addButton() {
  injectStyles();

  let url = window.location.href;
  if (!url.startsWith("https://plugins.jetbrains.com/plugin/")) {
    return;
  }

  let pluginId = url.split("/")[4].split("-")[0];
  console.log("pluginId: " + pluginId);

  let pluginDetail = await fetch(
    "https://plugins.jetbrains.com/api/plugins/" + pluginId,
  ).then((r) => r.json());

  const parentElement = await findElementWithRetry(
    ".plugin-header__controls-panel > div:first-child",
  );

  if (parentElement.querySelector(".jetbra-button")) {
    return;
  }
  let newElement = document.createElement("div");
  newElement.classList.toggle("wt-col-inline");
  newElement.innerHTML = `<div class="customize-btn" style="display: flex; flex-direction: column;white-space: nowrap;">
          <div class="generate-plugin-code" style="display: flex;height: 40px;">
              <button class="jetbra-button" type="button" id="permanentLicense">获取永久激活码</button>
              <button class="jetbra-button" type="button" id="twoYearsLicense">获取一年激活码</button>
          </div>
      </div>`;

  parentElement.appendChild(newElement);

  newElement.addEventListener("click", async (e) => {
    let productCode = pluginDetail?.purchaseInfo?.productCode;
    if (productCode === undefined) {
      window.alert("该插件不是市场付费插件！");
      return;
    }
    let licenseId = genLicenseId();
    let licensePartJson = buildLicensePartJson(
      productCode,
      licenseId,
      e.target.id === "twoYearsLicense",
    );

    try {
      const privateKey = importPrivateKeyFromPem(pemEncodedKey);
      const licensePartBase64 = btoa(
        unescape(encodeURIComponent(licensePartJson)),
      );
      const sigResultsBase64 = signData(privateKey, licensePartJson);
      const cert_base64 = pem2base64(pemEncodedCrt);

      GM_setClipboard(
        `${licenseId}-${licensePartBase64}-${sigResultsBase64}-${cert_base64}`,
        "text",
      );
      window.alert("激活码已经拷贝到你的剪切板");
    } catch (error) {
      console.error("Cryptographic operation failed:", error);
    }
  });
}

addButton();
if (window.onurlchange === null) {
  window.addEventListener("urlchange", (ignore) => {
    addButton();
  });
}
