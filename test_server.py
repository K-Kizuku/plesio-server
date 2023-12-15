import socket
import sys

server_addr = ('', 8254)

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

sock.bind(server_addr)

M_SIZE = 1024

while True:
    print("wait_for_msg", file=sys.stderr)
    bytes_message, client_addr = sock.recvfrom(M_SIZE)
    decoded_message = bytes_message.decode()
    print(decoded_message, file=sys.stderr)
    send_message = decoded_message + "recieved. v6"
    i = 0
    while i < 100000000:
        i = i + 1;

    send_message = send_message + str(i)

    sock.sendto(send_message.encode(), client_addr)
