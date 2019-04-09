# Google Maps API


      function initMap() {
        const map = new google.maps.Map(document.getElementById('map'), {
          center: {lat: 48.85270700, lng: 2.41379600},
          zoom: 17
        });
        const marker = new google.maps.Marker({
            position: {
                lat: 42.32,
                lng: 2.41
            },
            map: map, //La carte sur laquelle placer le marqueur
            icon: "voir la documentation"
        });

        const infoWindow = new google.maps.InfoWindow({
            content: "HTML" //Il faut ajouter un event listener
        });

        marker.addListener("click", () => infoWinfow.open(map, marker))
      }

      Pour ajouter plusieurs marqueurs :
      // Add marker function
      addMarker({

      })

      const addMarker = coords => {
            const marker = new google.maps.Marker({
            position: props.coords,
            map: map, //La carte sur laquelle placer le marqueur
            icon: props.iconImage
        })
        if(props.iconImage)
            marker.setIcon(props.iconImage)
        ;
      }

      Placer un marqueur de manière interactivie (directement sur la carte)

      //Listen for click on map
      google.maps.event.addListener(map, 'click', event => {
          addMarker({coords: event.latLng});
      })