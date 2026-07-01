# Hướng Dẫn Truy Vấn Grafana (LogsQL)

## Cách Truy cập Grafana
1. Mở trình duyệt web và truy cập vào `http://localhost:3000`.
2. Bấm "Sign in with Google" (Nếu đã thiết lập SSO) HOẶC đăng nhập bằng tài khoản mặc định:
   - **Tên đăng nhập**: `admin`
   - **Mật khẩu**: `admin`
3. Ở lần đăng nhập đầu tiên (nếu dùng tài khoản admin), bạn có thể được yêu cầu đổi mật khẩu. Bạn có thể bỏ qua (Skip) hoặc đặt mật khẩu mới.
4. Chuyển đến tab **Explore** (biểu tượng la bàn ở thanh menu bên trái).
5. Chọn nguồn dữ liệu **VictoriaLogs** từ menu thả xuống ở trên cùng.

## Các Ví dụ về LogsQL
Ứng dụng Golang tự động sinh log với các nhãn (labels) là `{job="golang-payment-service", env="dev"}`. Dưới đây là một số câu lệnh truy vấn LogsQL thực tế để khám phá dữ liệu.

### 1. Lọc theo Tên dịch vụ (Job)
Để xem toàn bộ log được sinh ra bởi dịch vụ payment của chúng ta:
```logsql
_stream:{job="golang-payment-service"}
```
### 2. Lọc theo Mức độ Lỗi (Error Level)
Để tìm kiếm tất cả các log báo lỗi (error) từ môi trường dev:
```logsql
_stream:{env="dev"} AND level:"error"
```

### 3. Tìm kiếm Toàn văn bản (Full-Text Search)
Để tìm kiếm một giao dịch cụ thể hoặc hành động của người dùng trong toàn bộ log:
```logsql
_stream:{job="golang-payment-service"} AND "Transaction failed"
```

*Mẹo: Bạn có thể sử dụng nút chọn `Time range` (Khoảng thời gian) ở góc trên bên phải của màn hình Explore để thu hẹp kết quả tìm kiếm trong 5 phút, 1 giờ qua, v.v.*
