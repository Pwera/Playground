#include <fplus/fplus.hpp>
#include <iostream>
#include <cpr/cpr.h>
#include "oatpp/core/async/Executor.hpp"
API_CLIENT_INIT(DemoApiClient)
class AsyncExample {
private:
    typedef oatpp::web::protocol::http::incoming::Response Response;
private:
    constexpr static const char* TAG = "AsyncExample";
private:

    class SendDtoCoroutine : public oatpp::async::Coroutine<SendDtoCoroutine> {
    private:
        std::shared_ptr<DemoApiClient> m_client;
        oatpp::String m_message;
        v_int32 m_code;
    public:

        SendDtoCoroutine(const std::shared_ptr<DemoApiClient> client,
                         const oatpp::String& message,
                         v_int32 code)
                : m_client(client)
                , m_message(message)
                , m_code(code)
        {}

        /**
         * Create DTO and send it via doPostWithDtoAsync method
         */
        Action act() override {
            auto dto = MyRequestDto::createShared();
            dto->message = m_message;
            dto->code = m_code;
            return m_client->doPostWithDtoAsync(dto).callbackTo(&SendDtoCoroutine::onResponse);
        }

        /**
         * Receive response and read body asynchronously
         */
        Action onResponse(const std::shared_ptr<Response>& response) {
            return response->readBodyToStringAsync().callbackTo(&SendDtoCoroutine::onBody);
        }

        /**
         * Print read body from the stream and finish
         */
        Action onBody(const oatpp::String& body) {
            OATPP_LOGD(TAG, "[doPostWithDtoAsync] data='%s'", body->c_str());
            return finish();
        }

    };

    class SendCoroutine : public oatpp::async::Coroutine<SendCoroutine> {
    private:
        std::shared_ptr<DemoApiClient> m_client;
    public:

        SendCoroutine(const std::shared_ptr<DemoApiClient> client) : m_client(client) {}

        Action act() override {
            return m_client->doPostAsync("<POST-DATA-HERE>").callbackTo(&SendDtoCoroutine::onResponse);
        }

        Action onResponse(const std::shared_ptr<Response>& response) {
            return response->readBodyToStringAsync().callbackTo(&SendDtoCoroutine::onBody);
        }

        Action onBody(const oatpp::String& body) {
            OATPP_LOGD(TAG, "[SendCoroutine. doPostAsync] data='%s'", body->c_str());
            return finish();
        }

    };

public:

    void static runExample(const std::shared_ptr<DemoApiClient>& client);

};
int main() {
    if (true) {

        std::list<std::string> things = {"same old", "same old"};
        if (fplus::all_the_same(things)) {
            std::cout << "All things being equal." << std::endl;
        }
    }

//    cpr::Response r = cpr::Get(cpr::Url{"http://api.github.com/repos/whoshuu/cpr/contributors"});
//                               cpr::Authentication{"user", "pass", AuthMode::BASIC},
//                               cpr::Parameters{{"anon", "true"}, {"key", "value"}}
//
//    r.status_code;                  // 200
//    r.header["content-type"];       // application/json; charset=utf-8
//    r.text;                         // JSON text string
//    std::cout << r.text << " " << std::endl;
    oatpp::async::Executor executor(1, 1, 1);
}