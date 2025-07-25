# 1) Build aşaması
FROM node:16-alpine AS build

# Çalışma dizini
WORKDIR /app

# Sadece package.json ve package-lock.json'ı kopyala
COPY package*.json ./

# Lockfile üzerinden temiz kurulum yap
RUN npm ci

# Uygulama kodunu kopyala ve derle
COPY . .
RUN npm run build

# 2) Serve aşaması
FROM nginx:alpine

# Özel Nginx konfigürasyonunu ekle (SPA history mode ve /api proxy için)
COPY nginx/default.conf /etc/nginx/conf.d/default.conf

# Build çıktısını Nginx’in HTML klasörüne kopyala
COPY --from=build /app/dist /usr/share/nginx/html

# 80 portunu expose et
EXPOSE 80

# Nginx’i ön planda çalıştır
CMD ["nginx", "-g", "daemon off;"]
