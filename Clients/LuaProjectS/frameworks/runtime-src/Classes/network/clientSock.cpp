#include "clientSock.h"
#include "errno.h"
#include "assert.h"
#include <string.h>
#include <thread>
#include "cocos2d.h"

#ifndef WIN32
#include <unistd.h>
#include "netdb.h"
#include <sys/ioctl.h>
#include "sys/types.h"
#include "sys/socket.h"
#else
#define ioctl ioctlsocket
#endif

#include "stdio.h"

Packet* createNetworkPacket(ClientStatus status, unsigned int ptid = 0, char* content = 0, unsigned short len = 0) {
	Packet* p = (Packet*)malloc(sizeof(Packet));

	p->status = status;
	p->ptid = ptid;
	p->content = content;
	p->len = len;

	return p;
}

void deleteNetworkPacket(Packet* p) {
	free((void *)p->content);
	free(p);
};


void setKeepAlive(int s)
{
	struct timeval tv;
	tv.tv_sec = 10;
	tv.tv_usec = 100000; // 0.1 sec
	setsockopt(s, SOL_SOCKET, SO_SNDTIMEO, (const char*)&tv, sizeof(tv));

#ifdef WIN32
	int keepalive = 1;
	int keepidle = 1;
	setsockopt(s, SOL_SOCKET, SO_KEEPALIVE, (const char*)&keepalive, sizeof(keepalive));

	tcp_keepalive alive_in = { 0 };
	tcp_keepalive alive_out = { 0 };
	alive_in.keepalivetime = 1000;
	alive_in.keepaliveinterval = 1000;
	alive_in.onoff = true;
	unsigned long ulBytesReturn = 0;

	WSAIoctl(s, SIO_KEEPALIVE_VALS, &alive_in, sizeof(alive_in), &alive_out, sizeof(alive_out), &ulBytesReturn, NULL, NULL);
#else
	int keepAlive = 1;
	int keepIdle = 1;
	int keepInterval = 1;
	int keepCount = 5;

	setsockopt(s, SOL_SOCKET, SO_KEEPALIVE, (void*)&keepAlive, sizeof(keepAlive));
#if defined(__APPLE__)
	setsockopt(s, IPPROTO_TCP, TCP_KEEPALIVE, (void*)&keepIdle, sizeof(keepIdle));
	setsockopt(s, IPPROTO_TCP, TCP_KEEPINTVL, (void*)&keepInterval, sizeof(keepInterval));
	setsockopt(s, IPPROTO_TCP, TCP_KEEPCNT, (void*)&keepCount, sizeof(keepCount));
#else
	setsockopt(s, SOL_TCP, TCP_KEEPIDLE, (void*)&keepIdle, sizeof(keepIdle));
	setsockopt(s, SOL_TCP, TCP_KEEPINTVL, (void*)&keepInterval, sizeof(keepInterval));
	setsockopt(s, SOL_TCP, TCP_KEEPCNT, (void*)&keepCount, sizeof(keepCount));
#endif
#endif
}

char* clientSock::getIpByDomain(const char* domain)
{
	hostent* ht = gethostbyname(domain);
	char* pszIp = (char*)inet_ntoa(*(struct in_addr*)(ht->h_addr));

	return pszIp;
}

int clientSock::recv_len(char* buffer, int len)
{
	int ret = 0;
	int needLen = len;
	int totalLen = 0;

	while (needLen > totalLen)
	{
		ret = recv(m_sock, &buffer[totalLen], needLen, 0);
		if (ret <= 0)
		{
			printf("[NET] recv error: %s\n", strerror(errno));
			break;
		}
		totalLen += ret;
	}
	return ret;
}

int clientSock::send_len(char* buffer, int len)
{
	int ret = 0;
	int needLen = len;
	int totalLen = 0;

	while (needLen > totalLen)
	{
		ret = send(m_sock, &buffer[totalLen], needLen, 0);
		if (ret <= 0)
		{
			printf("[NET] send error: %s\n", strerror(errno));
			break;
		}
		totalLen += ret;
	}
	return ret;
}

void clientSock::connectByIp(const char* ip, int port, std::function<void(Packet* ret)> reply)
{
	memset(m_ip, 0, 32);
	memcpy(m_ip, ip, strlen(ip));
	m_port = port;
	m_onReply = reply;

	std::thread start_thread(&clientSock::start_thread, this);
	start_thread.detach();
}

void clientSock::connectByDomain(const char* domain, int port, std::function<void(Packet* ret)> reply)
{
	const char* ip = getIpByDomain(domain);
	connectByIp(ip, port, reply);
}

void clientSock::push2QueueSend(Packet* p)
{
	std::lock_guard<std::mutex> lock(mutexSend);

	queueSend.push(p);
}

void clientSock::cleanQueueSend()
{
	std::lock_guard<std::mutex> lock(mutexSend);

	while (!queueSend.empty())
	{
		Packet* packet = queueSend.front();
		queueSend.pop();

		deleteNetworkPacket(packet);
	}
}

void clientSock::push2QueueRecv(Packet* p)
{
	std::lock_guard<std::mutex> lock(mutexRecv);

	queueRecv.push(p);
}

void clientSock::cleanQueueRecv()
{
	std::lock_guard<std::mutex> lock(mutexRecv);

	while (!queueRecv.empty())
	{
		Packet* packet = queueRecv.front();
		queueRecv.pop();

		deleteNetworkPacket(packet);
	}
}

void clientSock::start_thread()
{
	std::lock_guard<std::mutex> lock(mutexStatus);
		
	do
	{
		struct addrinfo *answer, hint;
		memset(&hint, 0, sizeof(hint));
		hint.ai_family = AF_UNSPEC;
		hint.ai_socktype = SOCK_STREAM;

		int ret = getaddrinfo(m_ip, NULL, &hint, &answer);
		if (ret != 0) {
			printf("[NET] getaddrinfo error:%s\r\n", strerror(errno));
			break;
		}

		int ipv_family = 0;
		for (struct addrinfo *curr = answer; curr != NULL; curr = curr->ai_next) {
			if (curr->ai_family == AF_INET)
			{
				ipv_family = AF_INET;
			}
			else if (curr->ai_family == AF_INET6)
			{
				memset(m_ip, 0, sizeof(m_ip));

				struct sockaddr_in6* addr6 = (struct sockaddr_in6*)curr->ai_addr;
				inet_ntop(AF_INET6, &addr6->sin6_addr, m_ip, sizeof(m_ip));

				ipv_family = AF_INET6;
			}
		}
		freeaddrinfo(answer);

		unsigned long ul = 1;

		if (ipv_family == AF_INET) {
			//ipv4
			m_sock = socket(AF_INET, SOCK_STREAM, 0);
			if (m_sock < 0) {
				printf("[NET] create socket error:%s\r\n", strerror(errno));
				break;
			}

			setKeepAlive(m_sock);

			struct sockaddr_in serv_addr;
			serv_addr.sin_family = AF_INET;
			serv_addr.sin_port = htons(m_port);
			serv_addr.sin_addr.s_addr = inet_addr(m_ip);

			ioctl(m_sock, FIONBIO, &ul);

			ret = connect(m_sock, (struct sockaddr*)&serv_addr, sizeof(struct sockaddr_in));
		}
		else if (ipv_family == AF_INET6)
		{
			//ipv6
			m_sock = socket(AF_INET6, SOCK_STREAM, 0);
			if (m_sock < 0) {
				printf("[NET] create socket error:%s\r\n", strerror(errno));
				break;
			}

			setKeepAlive(m_sock);

			struct sockaddr_in6 serv_addr;
			memset(&serv_addr, 0, sizeof(serv_addr));
			serv_addr.sin6_family = AF_INET6;
			serv_addr.sin6_port = htons(m_port);
			if (inet_pton(AF_INET6, m_ip, &serv_addr.sin6_addr) < 0)
			{
				printf("[NET]ipv6 error addr!\n");
				break;
			}

			ioctl(m_sock, FIONBIO, &ul);

			ret = connect(m_sock, (struct sockaddr*)&serv_addr, sizeof(struct sockaddr_in6));
		}

		bool success = false;
		if (ret == -1)
		{
			struct timeval tm;
			tm.tv_sec = 5;
			tm.tv_usec = 0;
			fd_set set;
			FD_ZERO(&set);
			FD_SET(m_sock, &set);
			if (select(m_sock + 1, NULL, &set, NULL, &tm) > 0)
			{
				int error = 0;
				int len = sizeof(int);
				getsockopt(m_sock, SOL_SOCKET, SO_ERROR, (char*)&error, (socklen_t *)&len);
				if (0 == error) success = true;
			}
			else
				success = false;
		}

		ul = 0;
		ioctl(m_sock, FIONBIO, &ul);

		if (success) {
			printf("[NET] Connect Server Success [ip:%s:%d]!\n", m_ip, m_port);

			connected = true;

			push2QueueRecv(createNetworkPacket(ClientStatus::Succeed));

			std::thread send_thread(&clientSock::send_thread, this);
			send_thread.detach();

			std::thread recv_thread(&clientSock::recv_thread, this);
			recv_thread.join();

			cleanQueueRecv();
			cleanQueueSend();

			if (connected)
			{
				closeSocket();
				push2QueueRecv(createNetworkPacket(ClientStatus::Losted));
			}
			return;
		}
		else {
			printf("[NET] Connect Server Fail [ip:%s:%d]!\n", m_ip, m_port);
			closeSocket();
		}
	} while (0);

	push2QueueRecv(createNetworkPacket(ClientStatus::Failed));
}

void clientSock::send_thread()
{
	printf("[NET] Send thread start ...\n");

	bool succeed = true;

	std::queue<Packet*> queueWork;

	char bufferHead[HEAD_LEN];
	char bufferProtoc[PROTOC_LEN];

	while (connected && succeed)
	{
		{
			std::lock_guard<std::mutex> lock(mutexSend);

			queueWork.swap(queueSend);
		}

		while (!queueWork.empty())
		{
			Packet* packet = queueWork.front();
			queueWork.pop();

			do
			{
				if (succeed)
				{
					*(unsigned short*)bufferHead = htons(PROTOC_LEN + packet->len);
					if (send_len(bufferHead, HEAD_LEN) <= 0)
					{
						printf("[NET] Send head error: %s\n", strerror(errno));
						succeed = false;
						break;
					}

					*(unsigned int*)bufferProtoc = htonl(packet->ptid);
					if (send_len(bufferProtoc, PROTOC_LEN) <= 0)
					{
						printf("[NET] Send protoc error: %s\n", strerror(errno));
						succeed = false;
						break;
					}

					if (packet->len > 0 && send_len(packet->content, packet->len) <= 0)
					{
						printf("[NET] Send content error: %s\n", strerror(errno));
						succeed = false;
						break;
					}
				}
			} while (0);

			deleteNetworkPacket(packet);
		}
		std::this_thread::sleep_for(std::chrono::milliseconds(10));
	}

	printf("[NET] Send thread exit ...\n");
}

void clientSock::recv_thread()
{
	printf("[NET] Recv thread start ...\n");

	char bufferHead[HEAD_LEN];
	char bufferProtoc[PROTOC_LEN];
	while (true)
	{
		if (recv_len(bufferHead, HEAD_LEN) <= 0) {
			printf("[NET] Recv head error: %s\n", strerror(errno));
			break;
		}

		unsigned short needLen = ntohs(*(unsigned short*)bufferHead);
		if (needLen > 0)
		{
			if (recv_len(bufferProtoc, PROTOC_LEN) <= 0)
			{
				printf("[NET] Recv protoc error: %s\n", strerror(errno));
				break;
			}

			Packet* packet = (Packet*)malloc(sizeof(Packet));
			packet->status = ClientStatus::Dataed;
			packet->ptid = ntohl(*(unsigned int*)bufferProtoc);
			packet->content = NULL;
			packet->len = 0;

			int contentLen = needLen - PROTOC_LEN;
			if (contentLen > 0)
			{
				char* content = (char*)malloc(contentLen);
				if (recv_len(content, contentLen) <= 0)
				{
					printf("[NET] Recv content error: %s\n", strerror(errno));
					free(content);
					deleteNetworkPacket(packet);
					break;
				}
				
				packet->content = content;
				packet->len = contentLen;
			}
			push2QueueRecv(packet);
		}
	}

	printf("[NET] Recv thread exit ...\n");
}

void clientSock::closeSocket(){
	connected = false;

    shutdown(m_sock, 2);
#ifdef WIN32
	closesocket(m_sock);
#else
	close(m_sock);
#endif
}

void clientSock::checkUpdate()
{
	std::queue<Packet*> queueWork;
	{
		std::lock_guard<std::mutex> lock(mutexRecv);

		queueWork.swap(queueRecv);
	}

	while (!queueWork.empty())
	{
		Packet* packet = queueWork.front();
		queueWork.pop();

		m_onReply(packet);

		deleteNetworkPacket(packet);
	}
}

void clientSock::handlerSend(unsigned short ptid, char* data, unsigned short dataLen)
{
	Packet* packet = (Packet*)malloc(sizeof(Packet));
	packet->status = ClientStatus::Dataed;
	packet->ptid = ptid;
	packet->content = data;
	packet->len = dataLen;

	push2QueueSend(packet);
}
