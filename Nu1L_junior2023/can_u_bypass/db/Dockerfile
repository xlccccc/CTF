FROM absolutapps/oracle-12c-ee

ENV ORACLE_ALLOW_REMOTE=true
COPY flag /flag
COPY readflag /readflag
RUN chmod 0600 /flag && \
    chmod u+s /readflag && \
    chmod +x /readflag
