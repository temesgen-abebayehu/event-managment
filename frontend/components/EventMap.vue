<template>
  <div>
    <div id="map" class="w-full h-[500px] rounded-lg shadow-lg"></div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, watch } from 'vue'
import L from 'leaflet'

const props = defineProps<{
  events: any[]
  center?: [number, number]
}>()

const emit = defineEmits(['marker-click'])

let map: any = null
let markers: any[] = []

onMounted(() => {
  // Fix Leaflet default marker icon issue
  delete (L.Icon.Default.prototype as any)._getIconUrl
  L.Icon.Default.mergeOptions({
    iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
    iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
    shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
  })

  // Initialize map (default to Addis Ababa)
  const center = props.center || [9.005401, 38.763611]
  map = L.map('map').setView(center, 12)

  // Add OpenStreetMap tiles
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap contributors',
    maxZoom: 19,
  }).addTo(map)

  // Add markers
  updateMarkers()
})

watch(() => props.events, () => {
  updateMarkers()
}, { deep: true })

const updateMarkers = () => {
  if (!map) return

  // Clear existing markers
  markers.forEach(marker => marker.remove())
  markers = []

  // Add new markers
  props.events.forEach(event => {
    if (event.latitude && event.longitude) {
      const marker = L.marker([event.latitude, event.longitude])
        .addTo(map)
        .bindPopup(`
          <div class="text-center">
            <h3 class="font-semibold">${event.title}</h3>
            <p class="text-sm text-gray-600">${event.venue}</p>
            <p class="text-sm font-bold text-primary mt-1">
              ${event.price === 0 ? 'FREE' : event.price + ' ETB'}
            </p>
          </div>
        `)
        .on('click', () => {
          emit('marker-click', event)
        })

      markers.push(marker)
    }
  })

  // Fit map to show all markers
  if (markers.length > 0) {
    const group = L.featureGroup(markers)
    map.fitBounds(group.getBounds().pad(0.1))
  }
}

onBeforeUnmount(() => {
  if (map) {
    map.remove()
  }
})
</script>
