import socket
import sys
import io
import time
import picamera




# Create a UDS socket



#Open file
#myFile = open("shark.jpg",'rb') #open in binary 

# Connect the socket to the port where the server is listening
server_address = './socket'
print >>sys.stderr, 'connecting to %s' % server_address



i = 1
# Create an in-memory stream

with picamera.PiCamera() as camera:
    camera.start_preview()
    camera.resolution = (640, 480)
    camera.framerate = 30
    camera.vflip = True
    # Camera warm-up time
    time.sleep(2)
    myFile = io.BytesIO()
    for foo in camera.capture_continuous(myFile, 'jpeg',
                                             use_video_port=True):
        sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        sock.connect(server_address)
        connection = sock.makefile('wb')

        myFile.seek(0)
        print "Sending Binaries..."
        connection.write(myFile.read())
        myFile.seek(0)
        myFile.truncate()
        print "Success!!!"
        sock.close()











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

