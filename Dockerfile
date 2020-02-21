FROM scratch

# These commands copy your files into the specified directory in the image
# and set that as the working location
WORKDIR /textseg
COPY textseg.app /textseg/
COPY *.txt /textseg/

# This command runs your application, comment out this line to compile only
ENTRYPOINT ["./textseg.app"]

LABEL Name=patchyvideo-textseg Version=1.0.0
