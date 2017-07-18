FROM scratch
#accounting
EXPOSE 12881
#discovery
EXPOSE 10881
#menu
EXPOSE 12882
#game
EXPOSE 10887
#admin
EXPOSE 19770

COPY bin/pes6go /

ENTRYPOINT ["/pes6go"]

CMD ["fullhouse"]