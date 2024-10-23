FROM node:22-alpine as build

WORKDIR /webapp

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM node:22-alpine

WORKDIR /webapp

COPY --from=build /webapp/dist ./dist

RUN npm install -g serve

EXPOSE 8069

CMD ["serve", "-s", "dist", "-l", "8069"]
 
