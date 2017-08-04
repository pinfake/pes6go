call .\build-linux-binaries.bat
docker build -t pes6go .
for /F %%k in ('docker images -f "dangling=true" -q') do docker rmi %%k