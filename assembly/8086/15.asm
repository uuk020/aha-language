assume cs:codesg,ds:datasg

datasg segment
    db 'BaSiC'
    db 'iNfOrMaTiOn'
datasg ends

codesg segment
    start: mov ax,datasg
           mov ds,ax
           mov bx, 0
           mov cx, 5
        s: mov al, [bx]
           and al, 11011111b
           mov [bx],al
           inc bx
           loop s
           
           mov bx, 0
           mov cx, 11
        s0:mov al, [bx+5]
           or al, 00100000b
           mov [bx+5],al
           inc bx
           loop s0

           mov ax, 4c00h
           int 21h
codesg ends
end start