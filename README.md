# Sec2Crypto

De applicatie is geschreven in Golang. Er wordt gebruik gemaakt van de in de taal ingebouwde webserver om te de pagina te hosten. 
De executable binary is meegecommit zodat er niks gecompiled hoeft te worden om te testen. De server draait op port: 9090

De standaard crypto packages van golang hebben support voor AES encryption, deze zijn dus ook gebruikt om de data te versleutelen.

  <b>Versleutelen</b>
  Eerst wordt het wachtwoord dat de gebruiker geeft gehashed naar een key met MD5 (Dit zou veiliger zijn met SHA256)
  Dan wordt het bericht versleuteld en opgeslagen
  
  <b>Ophalen</b>
  Het bericht wordt opgehaald aan de hand van de gebruikers naam. 
  De key wordt weer gegenegeerd met MD5 en de content wordt ontcijferd.
  
De data wordt opgeslagen met Scribble, een kleine JSON database. 
 
