import io
import time
import picamera

# Create an in-memory stream
my_stream = io.BytesIO()
with picamera.PiCamera() as camera:
    camera.start_preview()
    camera.use_video_port = True
    # Camera warm-up time
    time.sleep(2)
    camera.capture(my_stream, 'jpeg')