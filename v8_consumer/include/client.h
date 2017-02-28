#ifndef CLIENT_H
#define CLIENT_H

#ifndef STANDALONE_BUILD
extern void assert(int);
#else
#include <cassert>
#endif

#include "commands.h"
#include "message.h"
#include "v8worker.h"

#include "../../flatbuf/include/header_generated.h"
#include "../../flatbuf/include/payload_generated.h"

#include <fstream>
#include <queue>
#include <uv.h>
#include <vector>

const size_t MAX_BUF_SIZE = 65536;

const int HEADER_FRAGMENT_SIZE = 4; // uint32
const int PAYLOAD_FRAGMENT_SIZE = 4; // uint32

std::ofstream cinfo_out;
std::ofstream cerror_out;

typedef struct {
  uv_write_t req;
  uv_buf_t buf;
} write_req_t;

typedef struct message_s {
  std::string header;
  std::string payload;
} message_t;

typedef struct header_s {
  uint8_t event;
  uint8_t opcode;
  std::string metadata;
} header_t;


class AppWorker {
public:
  static AppWorker *GetAppWorker();
  void Init(const std::string& appname, const std::string& addr, int port);

  void OnConnect(uv_connect_t *conn, int status);
  void OnRead(uv_stream_t *stream, ssize_t nread, const uv_buf_t *buf);
  void OnWrite(uv_write_t *req, int status);

  void WriteMessage(Message *msg);

  void ParseValidChunk(uv_stream_t *stream, int nread, const char *buf);

  void RouteMessageWithoutResponse(header_t *parsed_header,
                                   message_t *parsed_message);
  std::string RouteMessageWithResponse(header_t *parsed_header,
                                       message_t *parsed_message);

  std::vector<char> *GetReadBuffer();

private:
  AppWorker();
  ~AppWorker();

  V8Worker *v8worker;

  uv_loop_t main_loop;
  uv_tcp_t tcp_sock;
  uv_connect_t conn;
  uv_stream_t *conn_handle;
  struct sockaddr_in server_sock;

  bool main_loop_running;
  std::string app_name;

  std::string next_message;

  std::vector<char> read_buffer;
  MessagePool outgoing_queue;
};

#endif