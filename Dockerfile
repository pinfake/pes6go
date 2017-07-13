FROM scratch

EXPOSE 12881
EXPOSE 10881
EXPOSE 10887

COPY bin/pes6go /

ENTRYPOINT ["/pes6go"]

CMD ["fullhouse"]