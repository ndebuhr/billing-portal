FROM logstash:7.5.2

RUN logstash-plugin install logstash-input-syslog
RUN logstash-plugin install logstash-filter-useragent
RUN logstash-plugin install logstash-filter-date
RUN logstash-plugin install logstash-filter-geoip
RUN logstash-plugin install logstash-patterns-core
RUN rm -f /usr/share/logstash/pipeline/logstash.conf
RUN rm -f /usr/share/logstash/config/logstash.yml
ADD pipeline/logstash.conf /usr/share/logstash/pipeline/logstash.conf
ADD config/logstash.yml /usr/share/logstash/config/logstash.yml