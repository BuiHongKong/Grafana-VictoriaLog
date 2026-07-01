# Hệ Thống Mô Phỏng Thu Thập Log Tập Trung

## Tổng quan
Dự án này cung cấp một môi trường mô phỏng thu thập log tập trung sử dụng Docker Compose. Nó được thiết kế để minh họa một kiến trúc tách biệt giữa việc thu nhận log, lưu trữ và trực quan hóa dữ liệu sử dụng Fluent Bit, VictoriaLogs và Grafana.

Để xem giải thích trực quan về kiến trúc này, vui lòng mở: [Tài liệu Kiến trúc](file:///d:/PERSONAL/Grafana-VictoriaLog/docs/grafana-victoriaLog-fluentbit.html).

## Ngăn xếp Công nghệ (Tech Stack)
- **Hạ tầng (Infrastructure)**: Docker Compose
- **Gom Log (Log Aggregator/Gateway)**: Fluent Bit
- **Lưu trữ Log**: VictoriaLogs
- **Trực quan hóa (Visualization)**: Grafana OSS với plugin `victoriametrics-logs-datasource`
- **Ứng dụng Giả lập (Mock App)**: Golang Log Generator

## Kiến trúc Hệ thống & Luồng Dữ liệu
Môi trường này mô phỏng một phương pháp tiếp cận Endpoint (Gateway) tiêu chuẩn doanh nghiệp, phù hợp cho các máy chủ nội bộ:

1. **Golang App (Ứng dụng giả lập)**: Mô phỏng một dịch vụ nội bộ hoặc bên ngoài. Nó tự động sinh ra các log giả lập và gửi chúng qua mạng bằng phương thức HTTP POST đến cổng (endpoint) của Fluent Bit.
2. **Fluent Bit (Endpoint)**: Đóng vai trò là một cổng lắng nghe tập trung ở port `8888`. Nó tiếp nhận các log định dạng HTTP JSON, xử lý chúng, và đẩy tiếp xuống lớp lưu trữ.
3. **VictoriaLogs (Lưu trữ)**: Nhận các log đã được xử lý từ Fluent Bit và lưu trữ chúng xuống ổ cứng một cách tối ưu.
4. **Grafana (Trực quan hóa)**: Truy vấn VictoriaLogs bằng ngôn ngữ LogsQL để hiển thị log theo thời gian thực trên các Dashboard đã được cấu hình sẵn.

## Hướng dẫn Sử dụng
Khởi động toàn bộ hệ thống:
```bash
docker compose down
docker compose up -d --build
```
Truy cập Grafana tại `http://localhost:3000`. Các log giả lập sẽ tự động chảy vào dashboard "Golang Application Logs".
