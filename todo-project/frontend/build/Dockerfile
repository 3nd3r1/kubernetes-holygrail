FROM node:18-alpine AS base

FROM base AS build-stage
WORKDIR /usr/src/app

RUN apk add --no-cache libc6-compat

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM base AS run-stage
WORKDIR /usr/src/app

RUN adduser -S appuser && mkdir .next && chown -R appuser /usr/src/app

COPY --from=build-stage /usr/src/app/.next/standalone ./
COPY --from=build-stage /usr/src/app/.next/static ./.next/static

USER appuser

ENV PORT 3000
ENV HOSTNAME 0.0.0.0
EXPOSE ${PORT}

ENTRYPOINT ["node", "server.js"]
