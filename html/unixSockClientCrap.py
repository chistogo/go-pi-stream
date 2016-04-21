import io
import socket
import struct
import time
import picamera
import sys



# Create a UDS socket


print "Python Client"
#Open file
#myFile = open("shark.jpg",'rb') #open in binary 

# Connect the socket to the port where the server is listening
server_address = './socket'
print >>sys.stderr, 'connecting to %s' % server_address



sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)


sock.connect(server_address)

connection = sock.makefile('wb')


try:
    with picamera.PiCamera() as camera:
        camera.resolution = (640, 480)
        camera.framerate = 30
        time.sleep(2)
        start = time.time()
        stream = io.BytesIO()
        # Use the video-port for captures...
        for foo in camera.capture_continuous(stream, 'jpeg',
                                             use_video_port=True):
            connection.write(struct.pack('<L', stream.tell()))
            connection.flush()
            stream.seek(0)
            connection.write(stream.read())
            if time.time() - start > 30:
                break
            stream.seek(0)
            stream.truncate()
    connection.write(struct.pack('<L', 0))
finally:
    connection.close()
    client_socket.close()





# Create an in-memory stream

# with picamera.PiCamera() as camera:
#     camera.start_preview()
#     camera.resolution = (320, 240)
#     camera.use_video_port=True
#     camera.vflip = True
#     # Camera warm-up time
#     time.sleep(2)
    
#     while(True):
        
#         myFile = io.BytesIO()
#         sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
#         sock.connect(server_address)
#         camera.capture(myFile, 'jpeg')
#         myFile.seek(0)
#         print "Sending Binaries..."
#         l = myFile.read(1024)
#         while (l):
#             sock.send(l)
#             l = myFile.read(1024)
#         print "Success!!!"
#         sock.close()











# try:
    
#     # Send data
#     print "Sending Binaries..."
#     l = myFile.read(1024)
#     while (l):
#         sock.send(l)
#         l = myFile.read(1024)
#     print "Success!!!"

    
    
#     # while amount_received < amount_expected:
#     #     data = sock.recv(16)
#     #     amount_received += len(data)
#     #     print >>sys.stderr, 'received "%s"' % data

# finally:
#     print >>sys.stderr, 'closing socket'
#     sock.close()

