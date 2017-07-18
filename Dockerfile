FROM scratch

EXPOSE 12881 #accounting
EXPOSE 10881 #discovery
EXPOSE 12882 #menu
EXPOSE 10887 #game
EXPOSE 19770

COPY bin/pes6go /

ENTRYPOINT ["/pes6go"]

CMD ["fullhouse"]