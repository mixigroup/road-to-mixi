FROM ruby:3.0.2

LABEL maintainer="tomohiro.koike@mixi.co.jp"

WORKDIR /app

COPY Gemfile /app/Gemfile
RUN bundle install

COPY . /app
#EXPOSE 4567

CMD ["bundle", "exec", "unicorn", "-c", "unicorn.rb"]