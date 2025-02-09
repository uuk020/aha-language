assume cs:code
code segment
   mov ax, 0020h
   mov ds, ax

   mov dx, 0
   mov bx, 0
   mov cx, 64

s: mov [bx], dl
   inc bx
   inc dl

   loop s

   mov ax, 4c00h
   int 21h
code ends
end

