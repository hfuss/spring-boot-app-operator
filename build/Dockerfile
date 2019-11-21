FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/spring-boot-app-operator \
    USER_UID=1001 \
    USER_NAME=spring-boot-app-operator

# install operator binary
COPY build/_output/bin/spring-boot-app-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
