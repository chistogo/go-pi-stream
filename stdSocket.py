import socket
import sys
import io
import time
import picamera

HOST = 'wkuacm.org'
PORT = 1776

# Create an in-memory stream
with picamera.PiCamera() as camera:
    camera.start_preview()
    camera.resolution = (160, 120)
    camera.framerate = 60
    camera.vflip = True
    # Camera warm-up time
    time.sleep(2)
    myFile = io.BytesIO()
    for foo in camera.capture_continuous(myFile, 'jpeg', use_video_port=True):
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        s.connect((HOST,PORT))
        connection = s.makefile('wb')

        myFile.seek(0)
        print "Sending Binaries..."
        connection.write(myFile.read())
        myFile.seek(0)
        myFile.truncate()
        print "Success!!!"

        s.close()   
