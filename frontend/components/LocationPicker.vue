<template>
  <div>
    <label class="block text-sm font-medium text-gray-700 mb-2">
      Location * <span class="text-gray-500 text-xs">(Click map to set location)</span>
    </label>
    
    <!-- Map Container -->
    <div ref="mapContainer" class="h-96 rounded-lg overflow-hidden border border-gray-300 mb-2"></div>
    
    <!-- Coordinates Display -->
    <div class="grid grid-cols-2 gap-4 text-sm">
      <div>
        <span class="text-gray-600">Latitude:</span>
        <span class="ml-2 font-mono">{{ latitude.toFixed(6) }}</span>
      </div>
      <div>
        <span class="text-gray-600">Longitude:</span>
        <span class="ml-2 font-mono">{{ longitude.toFixed(6) }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import L from 'leaflet'
import 'leaflet/dist/leaflet.css'

const props = defineProps<{
  initialLat?: number
  initialLng?: number
}>()

const emit = defineEmits<{
  update: [{ lat: number; lng: number }]
}>()

const mapContainer = ref<HTMLElement | null>(null)
const latitude = ref(props.initialLat || 9.005401) // Addis Ababa default
const longitude = ref(props.initialLng || 38.763611)

let map: L.Map | null = null
let marker: L.Marker | null = null

onMounted(() => {
  if (!mapContainer.value) return

  // Fix Leaflet default marker icon
  delete (L.Icon.Default.prototype as any)._getIconUrl
  L.Icon.Default.mergeOptions({
    iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
    iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
    shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
  })

  // Initialize map
  map = L.map(mapContainer.value).setView([latitude.value, longitude.value], 13)

  // Add tile layer
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap contributors',
    maxZoom: 19,
  }).addTo(map)

  // Add initial marker
  marker = L.marker([latitude.value, longitude.value], { draggable: true }).addTo(map)

  // Handle marker drag
  marker.on('dragend', () => {
    if (marker) {
      const pos = marker.getLatLng()
      updateLocation(pos.lat, pos.lng)
    }
  })

  // Handle map click
  map.on('click', (e: L.LeafletMouseEvent) => {
    updateLocation(e.latlng.lat, e.latlng.lng)
    if (marker) {
      marker.setLatLng(e.latlng)
    }
  })
})

const updateLocation = (lat: number, lng: number) => {
  latitude.value = lat
  longitude.value = lng
  emit('update', { lat, lng })
}

onUnmounted(() => {
  if (map) {
    map.remove()
  }
})
</script>
