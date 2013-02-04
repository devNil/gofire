#ideen
hier werden alle ideen gesammelt.
##server
der server kombiniert einen http- und websocket-server. beide sind bereits in go implementiert. das front-end wird in html5 und javascript realisiert.
##chat
der chat könnte befehle unterstützen:

###das @

<strong>@USER MESSAGE</strong>

dem user USER wird eine nachricht MESSAGE gesendet

<strong>@ USER MESSAGE</strong>

das @ wird nicht interpretiert

<strong>@USER1, @USER2 MESSAGE</strong>

den user USER1 und USER2 wird eine nachricht MESSAGE gesendet.

<strong>@USER1, @USER2 MESSAGE1 $$ MESSAGE2</strong>

dem user USER1 wird die nachricht MESSAGE1 und dem user USER2 wird die nachricht MESSAGE2 geschickt
