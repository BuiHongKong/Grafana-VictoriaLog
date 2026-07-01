$ErrorActionPreference = 'Stop'

$GrafanaUrl = "http://localhost:3000"
# Thông tin xác thực mặc định của Grafana (admin:admin)
$EncodedCredentials = [System.Convert]::ToBase64String([System.Text.Encoding]::ASCII.GetBytes("admin:admin"))
$AuthHeader = @{ "Authorization" = "Basic $EncodedCredentials" }

Write-Host "============================================="
Write-Host " Cấu hình phân quyền Folder trong Grafana"
Write-Host "============================================="

# 1. Tạo Teams
$teams = @("project-a", "project-b", "project-c")
$teamIds = @{}

Write-Host "`n[1] Đang khởi tạo Teams..."
foreach ($team in $teams) {
    try {
        $search = Invoke-RestMethod -Uri "$GrafanaUrl/api/teams/search?name=$team" -Headers $AuthHeader -Method Get
        if ($search.teams.Count -eq 0) {
            $body = @{ name = $team } | ConvertTo-Json
            $res = Invoke-RestMethod -Uri "$GrafanaUrl/api/teams" -Method Post -Headers $AuthHeader -Body $body -ContentType "application/json"
            $teamIds[$team] = $res.teamId
            Write-Host "  [+] Đã tạo Team mới: $team (ID: $($res.teamId))"
        } else {
            $teamIds[$team] = $search.teams[0].id
            Write-Host "  [v] Team $team đã tồn tại (ID: $($search.teams[0].id))."
        }
    } catch {
        Write-Host "  [!] Lỗi khi xử lý Team $team : $_"
    }
}

# 2. Lấy danh sách thư mục (Folders)
Write-Host "`n[2] Lấy danh sách Thư mục (Folders)..."
$folders = $null
try {
    $folders = Invoke-RestMethod -Uri "$GrafanaUrl/api/folders" -Headers $AuthHeader -Method Get
} catch {
    Write-Host "[!] Không thể lấy danh sách Folders. Grafana đã chạy chưa?"
    exit
}

foreach ($folder in $folders) {
    # Chuyển tên folder "Project A" thành "project-a" để map với Team
    $folderTitle = $folder.title.ToLower().Replace(" ", "-")
    
    if ($teamIds.ContainsKey($folderTitle)) {
        $uid = $folder.uid
        $teamId = $teamIds[$folderTitle]
        
        Write-Host "  -> Cấu hình quyền cho Thư mục: '$($folder.title)'..."
        
        # Quyền 1 = View, 2 = Edit, 4 = Admin
        # Ghi đè toàn bộ quyền cũ, chỉ cho phép Team tương ứng được xem.
        $permissions = @(
            @{ teamId = $teamId; permission = 1 }
        )
        $body = @{ items = $permissions } | ConvertTo-Json -Depth 5
        
        try {
            Invoke-RestMethod -Uri "$GrafanaUrl/api/folders/$uid/permissions" -Method Post -Headers $AuthHeader -Body $body -ContentType "application/json"
            Write-Host "     [+] Đã gán Team '$folderTitle' vào Thư mục '$($folder.title)'"
        } catch {
            Write-Host "     [!] Lỗi khi gán quyền: $_"
        }
    }
}

Write-Host "`n[HOÀN TẤT] Bạn có thể gán user vào các Team này từ giao diện Grafana (Server Admin -> Teams)!"
