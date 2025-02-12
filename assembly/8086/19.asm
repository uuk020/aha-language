assume cs:codesg,ds:datasg

data segment
  db 'ibm             '
  db 'dec             '
  db 'dos             '
  db 'vax             '
  dw 0 ;定义一个字，用来暂存 cx
data ends

codesg segment
  start: mov ax,datasg
         mov ds,ax
         mov bx,0
         
         mov cx,4
      s0:mov ds:[40h],bx ; 将外层循环的 cx 的值保存在 datasg:40h(因为4个16字节的) 单元中
         mov si,0
         mov cx,3

       s:mov al,[bx+si]
         and al,11011111b
         mov [bx+si],al
         inc si
         loop s

         add bx,16
         mov cx,ds:[40h] ; 用 datasg:40h 单元中的 cx 的值恢复
         loop s0
         
         mov ax,4c00h
         int 21h
codesg ends

end start
