# Hệ Thống Gom Log Tập Trung (Log Aggregator)

> Mở Endpoint trên Server để thu thập log từ mạng nội bộ

## Flow Diagram

**[Các Nguồn Log]** ➔ **[Fluent Bit Endpoint]** ➔ **[VictoriaLogs]** ➔ **[Grafana]**

* **Các Nguồn Log:** Ứng dụng, Server khác bắn log qua mạng (HTTP/Forward) tới Server tập trung.
* **Fluent Bit Endpoint:** Mở port (Endpoint) trên Server. Lắng nghe, nhận log và chuẩn hóa.
* **VictoriaLogs:** Lưu trữ siêu tốc. Nhận dữ liệu từ Fluent Bit và ghi xuống ổ cứng.
* **Grafana:** Giao diện trực quan. Truy xuất vào VictoriaLogs để vẽ biểu đồ và tìm kiếm.

## Mô hình nhận Log qua Endpoint
Trong kiến trúc này, máy chủ nội bộ của bạn sẽ đóng vai trò như một **"Trạm thu gom tập trung" (Aggregator)**. Chúng ta sẽ cấu hình Fluent Bit mở ra một cổng lắng nghe (Endpoint mạng). Bất kỳ hệ thống hoặc ứng dụng nào từ bên ngoài chỉ cần bắn dữ liệu log thẳng vào địa chỉ IP và Port này.

- **Gom tụ mọi nơi:** Dù hệ thống của bạn phân tán ở hàng chục Server khác nhau trong mạng nội bộ, tất cả đều có thể cấu hình để bắn log về chung một máy chủ tập trung duy nhất.
- **Fluent Bit làm Cổng gác (Gateway):** Thay vì đọc file tĩnh, Fluent Bit sẽ mở HTTP Input hoặc Forward Input để hứng dữ liệu luân chuyển liên tục qua mạng.
- **Bảo mật:** Bạn toàn quyền kiểm soát việc ai được phép bắn log vào hệ thống bằng cách cấu hình Tường lửa (Firewall) của máy chủ nội bộ.

## Các bước triển khai thực tế
Để xây dựng trạm thu log này trên máy chủ nội bộ, quy trình sẽ như sau:

1. **Chuẩn bị máy chủ:** Chuẩn bị một Server (Linux/Windows), gán IP tĩnh nội bộ và đảm bảo đã cài đặt Docker.
2. **Deploy Stack:** Đẩy cụm dịch vụ gồm Grafana, VictoriaLogs và Fluent Bit lên máy chủ bằng Docker Compose (`docker compose up -d`).
3. **Mở Endpoint trên Fluent Bit:** Cấu hình file `fluent-bit.conf`, sử dụng block `[INPUT]` dạng HTTP hoặc Forward để mở một Port (ví dụ: Port `8888`). Port này sẽ làm nhiệm vụ đón lõng dữ liệu log.
4. **Cấu hình Tường lửa (Firewall):** Mở Port `8888` trên tường lửa của máy chủ để cho phép các Server khác gửi log tới, đồng thời mở Port `3000` cho giao diện Grafana.
5. **Hoàn thiện:** Mọi thứ đã sẵn sàng! Bạn chỉ cần cho các ứng dụng nội bộ bắn dữ liệu vào địa chỉ `http://[Server-IP]:8888`, và mở Grafana lên để theo dõi trực tiếp.

## Phân luồng Đa Dự Án & Môi Trường
Hệ thống hỗ trợ quản lý **Multi-Project** (Đa dự án) trên cùng một máy chủ duy nhất thông qua cơ chế Gắn nhãn (Labeling) và Truy vấn (Query):

- **Đánh nhãn từ nguồn:** Ứng dụng khi gửi log sẽ đính kèm thông tin dự án và môi trường dưới dạng chuỗi JSON (ví dụ: `"project": "Project A"`, `"env": "prod"`).
- **Chỉ mục hóa (Indexing):** Nhờ tham số `_stream_fields=project,env` cấu hình trên Fluent Bit, VictoriaLogs sẽ tự động lấy 2 trường này làm chỉ mục phân loại để tăng tốc độ tìm kiếm.
- **Dashboard as Code:** Bằng tính năng *Provisioning*, Grafana tự động sinh ra các thư mục độc lập (Project A, B, C). Trong mỗi thư mục là các Dashboard riêng cho `dev`, `staging`, `prod` sử dụng câu lệnh LogsQL (ví dụ: `project:"Project A" AND env:"dev"`).

## Màn hình Giám sát Tổng (Global Catch-all)
Bên cạnh việc xé nhỏ log cho từng dự án, kiến trúc còn bổ sung một **Global Dashboard** mang tính chất quản trị tập trung:

- **Catch-all Query:** Sử dụng câu lệnh Query tối giản `*` để bắt toàn bộ mọi dòng log đang chảy vào hệ thống theo thời gian thực.
- **Orphaned Logs (Log không phân loại):** Nếu một ứng dụng gửi log về bị thiếu sót thông tin (không đính kèm nhãn `project` hoặc `env`), các log này sẽ bị loại trừ khỏi các Dashboard con do không khớp điều kiện lọc. Tuy nhiên, lệnh `*` trên màn hình Global sẽ thu thập và hiển thị đầy đủ chúng. Nhờ vậy, đội ngũ quản trị hệ thống (SysAdmin) có thể kịp thời phát hiện các sai sót trong cấu hình từ phía ứng dụng mà không lo mất mát dữ liệu.
