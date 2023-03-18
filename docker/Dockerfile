# https://hub.docker.com/_/golang
FROM golang:1.20.2-alpine3.17

# Create working and notes directories.
RUN mkdir /root/atlas && mkdir /root/notes

# Set working directory.
WORKDIR /root/atlas

# Install packages.
RUN apk update && \
  apk upgrade && \
  apk add bash bash-completion make nano starship vim

# https://starship.rs/guide/#%F0%9F%9A%80-installation
RUN echo 'eval "$(starship init bash)"' >> ~/.bashrc

# https://starship.rs/config/#prompt
RUN mkdir -p ~/.config && \
  wget https://raw.githubusercontent.com/revett/dotfiles/main/starship.toml \
    -O ~/.config/starship.toml

# Disable the [battery] starship configuration as not supported by container.
RUN sed -i '/\[battery\]/,+1 d' ~/.config/starship.toml

# Create .atlas-notes.yml config
RUN echo "path: /root/notes" > /root/.atlas-notes.yml
