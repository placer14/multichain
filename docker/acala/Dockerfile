FROM ubuntu:xenial

RUN apt-get update && apt-get install --yes software-properties-common curl git clang
RUN apt-get install --yes --no-install-recommends build-essential

# Install Rust
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"

# Make sure submodule.recurse is set to true to make life with submodule easier.
RUN git config --global submodule.recurse true

# Clone repository
RUN git clone https://github.com/AcalaNetwork/Acala.git

RUN mv Acala /app

# Build
WORKDIR /app
RUN make init
RUN make build

COPY run.sh /root/
RUN chmod +x ~/run.sh

# rpc port
EXPOSE 9933
# ws port
EXPOSE 9944

ENTRYPOINT ["./root/run.sh"]
