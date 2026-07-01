# Cẩm Nang Vận Hành (Runbook)

## Yêu cầu Hệ thống
- Đã cài đặt và đang chạy Docker Desktop.
- Có kiến thức cơ bản về các câu lệnh Docker.

## Cách Khởi động Hệ thống
1. Mở terminal (CMD/PowerShell) và di chuyển đến thư mục gốc của dự án này.
2. Build và khởi động các container bằng Docker Compose:
   ```bash
   docker compose up -d --build
   ```
3. Chờ một lát để các dịch vụ khởi tạo. Grafana có thể mất khoảng một phút để tải xuống và cài đặt plugin cần thiết.

## Cách Kiểm tra Log của Container
Bạn có thể theo dõi log của từng dịch vụ để đảm bảo chúng đang hoạt động bình thường:

- **Xem tất cả log**:
  ```bash
  docker compose logs -f
  ```
- **Xem log của ứng dụng Golang** (hữu ích để xem các cục dữ liệu JSON được sinh ra):
  ```bash
  docker compose logs -f golang-app
  ```
- **Xem log của VictoriaLogs**:
  ```bash
  docker compose logs -f victorialogs
  ```
- **Xem log của Grafana**:
  ```bash
  docker compose logs -f grafana
  ```

## Cách Dừng & Dọn dẹp Hệ thống
Để dừng các dịch vụ và xóa container cùng mạng nội bộ (network) mặc định:
```bash
docker compose down
```

Để dừng các dịch vụ và XÓA LUÔN dữ liệu lưu trữ (thao tác này sẽ xóa sạch toàn bộ log đã lưu):
```bash
docker compose down -v
```
