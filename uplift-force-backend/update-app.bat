@echo off
chcp 65001 >nul
echo ==========================================
echo     Uplift Force Backend 部署脚本
echo ==========================================
echo.

echo 🔍 检查当前环境...
go env GOOS GOARCH

echo.
echo 🔨 设置Linux交叉编译环境...
set GOOS=linux
set GOARCH=amd64

echo 📦 编译Linux版本...
go build -o uplift-force-backend .
if %errorlevel% neq 0 (
    echo ❌ 编译失败！
    goto :restore_env
)
echo ✅ 编译成功！

echo.
echo 🛑 停止远程服务...
ssh root@45.32.67.85 "cd /opt/uplift-force-backend && ./stop.sh"

echo.
echo 📤 上传应用到服务器...
scp uplift-force-backend root@45.32.67.85:/opt/uplift-force-backend/
if %errorlevel% neq 0 (
    echo ❌ 上传失败！
    echo 尝试强制覆盖...
    ssh root@45.32.67.85 "rm -f /opt/uplift-force-backend/uplift-force-backend"
    scp uplift-force-backend root@45.32.67.85:/opt/uplift-force-backend/
    if %errorlevel% neq 0 (
        echo ❌ 强制上传也失败！
        goto :restore_env
    )
)
echo ✅ 上传成功！

echo.
echo 📤 上传配置文件...
if exist .env (
    scp .env root@45.32.67.85:/opt/uplift-force-backend/
    echo ✅ .env 文件已上传
) else (
    echo ⚠️  .env 文件不存在，跳过上传
)

echo.
echo 🔄 远程部署应用...
ssh root@45.32.67.85 "cd /opt/uplift-force-backend && ./deploy.sh"
if %errorlevel% neq 0 (
    echo ❌ 远程部署失败！
    goto :restore_env
)

echo.
echo 🌐 检查部署状态...
ssh root@45.32.67.85 "cd /opt/uplift-force-backend && ./status.sh"

echo.
echo ✅ 部署完成！
echo 📋 访问地址: http://45.32.67.85:8080
echo 📋 API测试: http://45.32.67.85:8080/api/v1/auth/register

:restore_env
echo.
echo 🔧 恢复Windows环境变量...
set GOOS=windows
set GOARCH=amd64
echo ✅ 环境已恢复为Windows

echo.
echo 🔍 当前环境:
go env GOOS GOARCH

echo.
echo 🧹 清理临时文件...
if exist uplift-force-backend (
    del uplift-force-backend
    echo ✅ 临时Linux可执行文件已删除
)

echo.
echo ==========================================
echo            部署流程完成！
echo ==========================================
pause