# THÔNG TIN

## Chức năng

- Quét qua N Byte dòng cuối cùng của file log
- Đếm số lượng dòng nào bắt đầu bằng `ERR`, `ERROR`, `EXCEPTION` --> M (mô tả thêm sau khi hoàn thành)
  -- Xác đinh bằng cách match dòng với regularexpress --> vd: `$YYYY-MM-DD HH:MM:SS.*(.exception.).*`
- Kiểm tra M <= [số lượng lỗi cho phép] hay không
  -- Yes --> exitcode = 0
  -- No --> exitcode = 1

## Sử dụng
`<tên app> [tên file log cần check] [số lượng BYTE cần quét] [Số lượng dòng ERR cho phép (bé hơn hoặc bằng số này)]`