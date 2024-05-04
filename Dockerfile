FROM imbios/bun-node:1.1.7-20.12.2-slim
WORKDIR /usr/app

COPY . .
RUN bun install && bun prisma generate

CMD [ "bun", "run", "start" ]