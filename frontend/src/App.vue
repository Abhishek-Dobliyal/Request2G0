<template>
  <div class="grid grid-cols-2">
    <div class="header-bg grid left-card">
      <div class="flex justify-center items-center">
        <div class="text-center text-white px-6 md:px-12">
          <h1 class="text-5xl font-bold mt-0 mb-6">{{ headerTitle }}</h1>
          <h3 class="text-2xl font-bold mb-8">{{ subHeaderTitle }}</h3>
        </div>
      </div>
    </div>
    <div class="grid h-screen">
      <div class="container flex mt-8 mx-5 h-1/2 items-center justify-center">
        <div class="dropdown relative">
          <button
            class="dropdown-toggle px-4 py-3 bg-black text-white font-medium text-xs leading-tight uppercase flex items-center whitespace-nowrap"
            type="button"
            id="dropdownMenuButton1"
            data-bs-toggle="dropdown"
            aria-expanded="false"
            ref="dropdownBtn"
          >
            Method
            <svg
              aria-hidden="true"
              focusable="false"
              data-prefix="fas"
              data-icon="caret-down"
              class="w-2 ml-2"
              role="img"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 320 512"
            >
              <path
                fill="currentColor"
                d="M31.3 192h257.3c17.8 0 26.7 21.5 14.1 34.1L174.1 354.8c-7.8 7.8-20.5 7.8-28.3 0L17.2 226.1C4.6 213.5 13.5 192 31.3 192z"
              ></path>
            </svg>
          </button>
          <ul
            class="dropdown-menu w-full absolute hidden bg-white text-base z-50 float-left py-2 list-none text-center shadow-lg hidden m-0 bg-clip-padding border-4"
            aria-labelledby="dropdownMenuButton1"
          >
            <li
              v-for="method in httpMethods"
              :key="method"
              @click="selectDropdownMethod"
            >
              <span
                class="dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-black hover:bg-gray-200 hover:cursor-hand"
                >{{ method }}</span
              >
            </li>
          </ul>
        </div>
        <form class="w-full max-w-sm px-4 mx-2">
          <div class="flex items-center border-b border-black">
            <input
              class="bg-transparent border-none w-full text-black mr-3 py-1 px-2 leading-tight focus:outline-none font-medium text-xs"
              type="url"
              placeholder="https://www.url.com"
              aria-label="url"
              v-model="requestObject.url"
              ref="urlInput"
              @input="validateInputURL"
            />
            <button
              class="flex-shrink-0 bg-black border-black text-sm text-white py-2.5 px-3 leading-tight"
              type="button"
              @click="sendRequest"
            >
              Request
            </button>
          </div>
        </form>
      </div>
      <div
        class="container flex justify-center content-center mt-5"
        v-if="!showReceivedResponse"
      >
        <ul
          class="nav nav-tabs flex flex-col md:flex-row flex-wrap list-none border-b-0 pl-0"
          id="tabs-tab"
          role="tablist"
        >
          <li class="nav-item px-2" role="presentation">
            <a
              href="#tabs-profile"
              class="tab-link block font-medium text-md leading-tight border-x-0 border-t-0 border-b-2 border-transparent px-6 py-3 my-2 hover:border-transparent hover:bg-gray-100"
              id="tabs-profile-tab"
              data-bs-toggle="pill"
              data-bs-target="#tabs-profile"
              role="tab"
              aria-controls="tabs-profile"
              aria-selected="false"
              ><i class="fa fa-pen px-2 items-center"></i>Custom Params</a
            >
          </li>
          <li class="nav-item px-2" role="presentation">
            <a
              href="#tabs-messages"
              class="tab-link block font-medium text-md leading-tight border-x-0 border-t-0 border-b-2 border-transparent px-6 py-3 my-2 hover:border-transparent hover:bg-gray-100 focus:border-transparent"
              id="tabs-messages-tab"
              data-bs-toggle="pill"
              data-bs-target="#tabs-messages"
              role="tab"
              aria-controls="tabs-messages"
              aria-selected="false"
              ><i class="fa fa-code px-2 items-center"></i>JSON</a
            >
          </li>
        </ul>
      </div>
      <div class="container flex justify-center" v-if="!showReceivedResponse">
        <div class="tab-content" id="tabs-tabContent">
          <div
            class="tab-pane fade"
            id="tabs-profile"
            role="tabpanel"
            aria-labelledby="tabs-profile-tab"
          >
            <form class="w-full max-w-sm" ref="customParamsInput">
              <div
                class="grid py-2"
                v-for="(item, key) in customParamsInputDivs"
                :key="key"
              >
                <div
                  class="container flex items-center border-b-2 border-black py-2"
                >
                  <input
                    class="appearance-none bg-transparent border-r-2 border-black w-full text-black mr-3 py-1 px-2 leading-tight focus:outline-none text-sm"
                    type="text"
                    :placeholder="`Key ${item.count}`"
                    aria-label="Key"
                    v-model="item.key"
                  />
                  <input
                    class="appearance-none bg-transparent border-none w-full text-black mr-3 py-1 px-2 leading-tight focus:outline-none text-sm"
                    type="text"
                    :placeholder="`Value ${item.count}`"
                    aria-label="Value"
                    v-model="item.value"
                  />
                  <button
                    class="flex-shrink-0 border-2 border-black text-black hover:text-black text-md py-1 px-2 hover:bg-gray-200"
                    type="button"
                    @click="
                      item.btnSymbol == '+'
                        ? addCustomParamInput()
                        : removeCustomParamInput()
                    "
                  >
                    {{ item.btnSymbol }}
                  </button>
                </div>
              </div>
            </form>
          </div>
          <div
            class="tab-pane fade mb-15 pb-5 text-center"
            id="tabs-messages"
            role="tabpanel"
            aria-labelledby="tabs-profile-tab"
          >
            <label
              for="message"
              class="block mb-2 text-sm font-semibold text-gray-900 dark:text-black"
              >Input JSON Here</label
            >
            <textarea
              id="message"
              rows="10"
              cols="50"
              v-model="requestObject.json"
              class="block p-2.5 w-full text-sm text-black bg-white rounded border-2 border-black bg-gray-100"
            ></textarea>
            <button
              class="px-6 py-2.5 mt-2 bg-black text-white font-medium text-xs leading-tight"
              type="button"
              id="dropdownMenuButton1"
              data-bs-toggle="dropdown"
              aria-expanded="false"
              @click="prettify"
            >
              Prettify
            </button>
          </div>
        </div>
      </div>
      <div class="container justify-center" v-else>
        <div class="flex justify-center items-center" v-if="showSpinner">
          <div
            class="animate-spin rounded-full h-20 w-20 border-b-4 border-black"
          ></div>
        </div>
        <div class="container text-center" v-else>
          <span class="font-semibold text-sm">Received Response</span>
          <textarea
            id="resp"
            rows="13"
            cols="50"
            v-model="response.body"
            class="block p-5 pb-1 mx-auto w-4/5 text-sm text-black bg-white rounded border-2 border-black bg-gray-100"
          ></textarea>
          <span class="font-semibold text-xs px-2"
            >Status Code: {{ response.statusCode }}
          </span>
          <button
            class="px-6 py-2.5 mx-3 mt-2 bg-black text-white font-medium text-xs leading-tight"
            type="button"
            id="dropdownMenuButton1"
            data-bs-toggle="dropdown"
            aria-expanded="false"
            @click="newRequest"
          >
            Request Another
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      headerTitle: "Request2GO",
      subHeaderTitle: "Calling & Testing APIs Made Simple",
      backendEnpointURL: "http://127.0.0.1:8000/fetch-response",
      httpMethods: ["GET", "POST", "PUT", "DELETE"],
      showReceivedResponse: false,
      showSpinner: false,
      customParamsInputDivs: {
        1: {
          btnSymbol: "+",
          count: 1,
          key: "",
          value: "",
        },
      },
      customParamsDivsCnt: 1,
      requestObject: {
        url: "",
        json: null,
        customParams: null,
        httpMethod: "",
      },
      response: {
        body: null,
        statusCode: null,
      },
    };
  },

  computed: {
    validateInputURL() {
      let input = this.$refs.urlInput;
      if (this.isValidURL(this.requestObject.url)) {
        input.style.color = "black";
        return true;
      }

      input.style.color = "red";
      return false;
    },
  },

  methods: {
    selectDropdownMethod(event) {
      var dropdownBtn = this.$refs.dropdownBtn;
      dropdownBtn.innerHTML = `${event.srcElement.innerText} <svg
              aria-hidden="true"
              focusable="false"
              data-prefix="fas"
              data-icon="caret-down"
              class="w-2 ml-2"
              role="img"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 320 512"
            >
              <path
                fill="currentColor"
                d="M31.3 192h257.3c17.8 0 26.7 21.5 14.1 34.1L174.1 354.8c-7.8 7.8-20.5 7.8-28.3 0L17.2 226.1C4.6 213.5 13.5 192 31.3 192z"
              ></path>
            </svg>`;
    },

    jsonToStr(json) {
      if (json == null) {
        return json;
      }

      if (typeof json == "string" && json != "") {
        var strToJson = JSON.parse(json);
        return JSON.stringify(strToJson, null, 3);
      }
      return JSON.stringify(json, null, 3);
    },

    prettify() {
      this.requestObject.json = this.jsonToStr(this.requestObject.json);
    },

    isValidURL(string) {
      var res = string.match(
        /(http(s)?:\/\/.)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)/g
      );
      return res !== null;
    },

    sendRequest() {
      this.prettify();
      this.showSpinner = true;
      var customParams = {};
      for (const item in this.customParamsInputDivs) {
        let customParamKey = this.customParamsInputDivs[item].key;
        if (
          customParamKey == "" ||
          this.customParamsInputDivs[item].value == ""
        ) {
          continue;
        }
        customParams[customParamKey] = this.customParamsInputDivs[item].value;
      }

      this.requestObject.httpMethod = this.$refs.dropdownBtn.innerText;
      if (this.requestObject.json == null) {
        this.requestObject.customParams = customParams;
      }

      fetch(this.backendEnpointURL, {
        method: "POST",
        body: this.requestObject,
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then((res) => {
          this.response.statusCode = res.status;
          return res.json();
        })
        .then((json) => {
          console.log(json);
          this.response.body = JSON.stringify(json, null, 3);
          this.showSpinner = false;
        });

      this.showReceivedResponse = true;
    },

    addCustomParamInput() {
      if (this.customParamsDivsCnt == 4) {
        return;
      }
      this.customParamsDivsCnt++;
      this.customParamsInputDivs[this.customParamsDivsCnt] = {
        btnSymbol: "-",
        count: this.customParamsDivsCnt,
        key: "",
        value: "",
      };
    },

    removeCustomParamInput() {
      var seen = false;
      this.customParamsDivsCnt--;
      for (const key in this.customParamsInputDivs) {
        if (this.customParamsInputDivs[key].btnSymbol == "-" && !seen) {
          delete this.customParamsInputDivs[key];
          seen = true;
        } else if (this.customParamsInputDivs[key].btnSymbol == "-") {
          this.customParamsInputDivs[key].count--;
        }
      }
    },

    newRequest() {
      this.requestObject = {
        url: "",
        json: null,
        customParams: null,
        httpMethod: "",
      };
      this.showReceivedResponse = false;
    },
  },

  created() {
    document.title = "Request2GO";
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Poppins&display=swap");

* {
  margin: 0;
  padding: 0;
  font-family: "Poppins", sans-serif;
}

.header-bg {
  background-color: black;
}

.left-card {
  color: black;
}

.dropdown-item {
  cursor: pointer;
}

.tab-link.active {
  color: black;
  border-color: black;
}

textarea {
  resize: none;
  outline: none;
}

input[type="url"] {
  caret-color: black;
}
</style>
