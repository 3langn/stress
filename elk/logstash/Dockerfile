FROM docker.elastic.co/logstash/logstash:8.4.3
RUN rm -f /usr/share/logstash/pipeline/logstash.conf
RUN curl -Lo /usr/share/logstash/logstash-core/lib/jars/postgresql.jar https://jdbc.postgresql.org/download/postgresql-42.2.20.jar
