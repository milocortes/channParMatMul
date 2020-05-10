# Uso de canales para multiplicación paralela de matrices

Un *canal* conecta un proceso remitente con un proceso receptor. Los canales son *tipados*, lo que significa que 
debemos declara el tipo de mensajes que pueden ser enviados en el canal. En lo que sigue, suponemos que los canales son síncronos.
En sistemas operativos, los canales son denominados *pipes*; los cuales permiten construir programas conectando un conjunto ya existente de programas.


Con canales podemos realizar *pipelined computation*. En este caso, usaremos canales para 
realizar una multiplicación paralela de matrices.
