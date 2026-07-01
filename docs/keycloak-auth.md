# Tích hợp Keycloak và Phân quyền (Authorization) trong Grafana

Tài liệu này giải thích vai trò của Keycloak và cách hệ thống phân quyền (Authorization) được thiết kế và áp dụng vào dự án Centralized Logging.

## 1. Keycloak là gì? Tại sao lại dùng nó?

Keycloak là một công cụ Quản lý Định danh (Identity and Access Management - IAM) mã nguồn mở rất phổ biến. 
Trong kiến trúc của chúng ta, Keycloak đóng vai trò là **Identity Provider (IdP)** để cung cấp tính năng **SSO (Single Sign-On)**.

### Giải quyết bài toán gì?
Nếu không có Keycloak, bạn sẽ phải tạo tài khoản người dùng trực tiếp bên trong Grafana. Khi hệ thống phình to (có thêm Kibana, Prometheus, hoặc các tool nội bộ khác), người dùng sẽ phải nhớ hàng chục tài khoản và mật khẩu khác nhau. 

Với Keycloak:
- Người dùng chỉ cần đăng nhập **một lần duy nhất** tại cổng Keycloak.
- Grafana (đóng vai trò Service Provider) sẽ tin tưởng Keycloak, cho phép người dùng vào hệ thống mà không cần gõ lại mật khẩu.

## 2. Authentication vs Authorization

- **Authentication (Xác thực):** Kiểm tra xem người dùng *có đúng là người họ nhận hay không* (Ví dụ: Đăng nhập đúng Username và Password). Việc này 100% do Keycloak xử lý.
- **Authorization (Phân quyền):** Sau khi đã xác thực xong, hệ thống cần kiểm tra xem người dùng *được phép làm gì, được xem những gì*. Trong dự án này, Authorization được phối hợp giữa Keycloak (quản lý Group) và Grafana (quản lý quyền trên Thư mục).

## 3. Kiến trúc Phân quyền Cách ly theo Dự án (Project-based Isolation)

Mục tiêu cốt lõi: **Nhân sự của Project nào thì chỉ được xem Dashboard của Project đó.** (Ví dụ: `user-a` chỉ xem được dữ liệu log của Project A).

### Cách thức hoạt động:

1. **Tại Keycloak:**
   - Chúng ta tạo ra các Groups tương ứng với các dự án: `/project-a`, `/project-b`, `/project-c` và nhóm `/admin`.
   - Các users được gán vào các Groups này (VD: `user-a` thuộc nhóm `/project-a`).

2. **Giao tiếp OAuth/OIDC:**
   - Khi `user-a` đăng nhập thành công, Keycloak sẽ cấp cho Grafana một Token (JSON Web Token - JWT).
   - Trong Token này có chứa thuộc tính (claim): `"groups": ["/project-a"]`.

3. **Tại Grafana (Role & Team Mapping):**
   - Grafana đọc thuộc tính `groups` từ Token.
   - Nhờ tham số `GF_AUTH_GENERIC_OAUTH_ROLE_ATTRIBUTE_PATH`, nếu user thuộc nhóm `/admin`, Grafana sẽ thăng cấp user đó thành `Admin` của hệ thống. Nếu không, user sẽ mang quyền `Viewer`.
   - Quan trọng hơn, Grafana hỗ trợ tính năng **Team Sync**. User mang Group `/project-a` sẽ tự động được gán vào Team có tên là `project-a` bên trong Grafana.

4. **Giới hạn Thư mục (Folder Permissions):**
   - Các Dashboard của chúng ta đã được chia thành các thư mục riêng (Project A, Project B, Project C).
   - Grafana được cấu hình để: 
     - Thư mục **Project A**: Chỉ cho phép Team `project-a` xem.
     - Thư mục **Project B**: Chỉ cho phép Team `project-b` xem.
     - Thư mục **Global**: Tất cả mọi người đều xem được.
   
Nhờ luồng (flow) này, dữ liệu log được bảo vệ một cách nghiêm ngặt từ đầu đến cuối mà không cần phải can thiệp thủ công mỗi khi có nhân sự mới gia nhập dự án!
