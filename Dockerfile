FROM scratch

# These commands copy your files into the specified directory in the image
# and set that as the working location
COPY textseg.app /textseg/
COPY *.txt /textseg/
WORKDIR /textseg

# This command runs your application, comment out this line to compile only
CMD ["./textseg.app"]

LABEL Name=patchyvideo-textseg Version=1.0.0
