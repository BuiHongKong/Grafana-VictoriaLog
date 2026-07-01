# Hướng dẫn Thiết lập Đăng nhập bằng Google (Google SSO) & Phân quyền

Hệ thống của chúng ta đã được tối ưu hóa để loại bỏ hoàn toàn các dịch vụ dư thừa (Keycloak, Postgres) và sử dụng trực tiếp tính năng **Đăng nhập bằng Google** được tích hợp sẵn bên trong Grafana.

## 1. Cách thiết lập Google Client ID và Secret

Để nút "Sign in with Google" hoạt động, bạn cần cung cấp cho Grafana mã định danh từ Google. Hãy làm theo các bước sau:

1. Truy cập vào **[Google Cloud Console](https://console.cloud.google.com/)**.
2. Tạo một Project mới (hoặc chọn Project hiện có).
3. Tìm kiếm và truy cập menu **APIs & Services > OAuth consent screen**.
   - Chọn loại User là **External** (hoặc Internal nếu bạn dùng Google Workspace của công ty).
   - Điền các thông tin bắt buộc (Tên App, Email hỗ trợ).
4. Chuyển sang menu **Credentials**.
   - Bấm **Create Credentials > OAuth client ID**.
   - Chọn Application type là **Web application**.
   - Tại mục **Authorized redirect URIs**, bạn bắt buộc phải thêm đường dẫn sau:
     `http://localhost:3000/login/google` (Nếu server của bạn chạy domain thật, hãy thay `http://localhost:3000` bằng domain thật của bạn, ví dụ: `https://grafana.congty.com/login/google`).
5. Bấm Create. Google sẽ cung cấp cho bạn 2 thông số:
   - **Client ID**
   - **Client Secret**

6. Mở file `docker-compose.yml`, tìm đến phần `grafana` và thay thế 2 mã này vào đúng chỗ của `<YOUR_GOOGLE_CLIENT_ID>` và `<YOUR_GOOGLE_CLIENT_SECRET>`.

## 2. Quy trình Phân quyền sau khi có Google SSO

Do Google không quản lý "Team Project A" của bạn, nên quy trình phân quyền sẽ được thực hiện 100% trên giao diện Grafana.

### Bước 1: User Đăng nhập lần đầu
Nhân sự của bạn mở Grafana, bấm nút "Sign in with Google" và đăng nhập bằng tài khoản Gmail thành công. Lúc này, User của họ đã được tự động tạo ngầm trong Grafana.

### Bước 2: Admin Cấp quyền
Bạn (với tư cách là Server Admin - đăng nhập bằng tài khoản `admin`/`admin` nếu chưa đổi) sẽ làm như sau:
1. Vào menu **Administration > Teams**.
2. Tìm đến `Team Project A` (hoặc tự tạo nếu chưa có).
3. Bấm **Add member** và chọn cái email Gmail vừa đăng nhập lúc nãy vào Team.
4. Xong! Do Thư mục Project A đã được cài đặt là "Chỉ Team Project A mới được xem", nên chủ nhân của Gmail đó từ giờ trở đi sẽ chỉ thấy được dữ liệu của Project A.

## 3. Quản lý Quyền nâng cao (Tùy chọn)
Nếu bạn muốn chặn tất cả Gmail lạ không cho phép đăng nhập (chỉ cho phép Gmail công ty), bạn có thể thêm biến môi trường sau vào file `docker-compose.yml`:
`- GF_AUTH_GOOGLE_ALLOWED_DOMAINS=tencongty.com`
