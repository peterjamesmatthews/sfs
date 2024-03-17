FROM node:21 AS base

FROM base as dependencies
WORKDIR /sfs
COPY package.json package-lock.json ./
RUN npm install

FROM dependencies as source
COPY . .

FROM source as debug
EXPOSE 5173
STOPSIGNAL SIGKILL
CMD npm run dev
