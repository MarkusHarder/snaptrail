FROM golang:1.24-bookworm

RUN useradd -ms /bin/bash gopher

COPY --chown=gopher:gopher ./be /home/gopher

COPY --chown=gopher:gopher ./web/dist /home/gopher/dist

ENV UI_DIR=/home/gopher/dist/ui/browser

RUN chmod +x /home/gopher/be

USER gopher

ENTRYPOINT ["/home/gopher/be"]
