<template>
  <div class="relative">
    <div ref="mapContainer" class="w-full h-[500px] rounded-lg shadow-lg bg-gray-100 relative overflow-hidden"></div>
    <div v-if="!tilesLoaded" class="absolute inset-0 flex items-center justify-center bg-gray-100 bg-opacity-90 pointer-events-none">
      <div class="text-center">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-purple-600 mb-2"></div>
        <p class="text-gray-600">Loading map...</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, watch } from 'vue'
import L from 'leaflet'
import 'leaflet/dist/leaflet.css'

const props = defineProps<{
  events: any[]
  center?: [number, number]
}>()

const emit = defineEmits(['marker-click'])

const mapContainer = ref<HTMLElement | null>(null)
const tilesLoaded = ref(false)
let map: L.Map | null = null
let markers: L.Marker[] = []

onMounted(() => {
  if (!mapContainer.value) return

  try {
    // Fix Leaflet default marker icon issue
    delete (L.Icon.Default.prototype as any)._getIconUrl
    L.Icon.Default.mergeOptions({
      iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
      iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
      shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
    })

    // Initialize map (default to Addis Ababa)
    const center = props.center || [9.005401, 38.763611]
    map = L.map(mapContainer.value).setView(center, 12)

    // Add OpenStreetMap tiles
    const tileLayer = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '© OpenStreetMap contributors',
      maxZoom: 19,
    }).addTo(map)

    // Listen for tile load events
    tileLayer.on('load', () => {
      tilesLoaded.value = true
    })

    // Add markers after a short delay to ensure map is ready
    setTimeout(() => {
      updateMarkers()
    }, 100)

    // Hide loading after 3 seconds even if tiles haven't loaded (timeout)
    setTimeout(() => {
      tilesLoaded.value = true
    }, 3000)
  } catch (error) {
    tilesLoaded.value = true
  }
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
      try {
        const marker = L.marker([event.latitude, event.longitude])
          .addTo(map!)
          .bindPopup(`
            <div class="text-center p-2">
              <h3 class="font-semibold text-base mb-1">${event.title}</h3>
              <p class="text-sm text-gray-600 mb-1">${event.venue}</p>
              <p class="text-sm font-bold text-purple-600">
                ${event.price === 0 ? 'FREE' : event.price + ' ETB'}
              </p>
            </div>
          `)
          .on('click', () => {
            emit('marker-click', event)
          })

        markers.push(marker)
      } catch (error) {
        // Silently skip invalid markers
      }
    }
  })

  // Fit map to show all markers
  if (markers.length > 0) {
    try {
      const group = L.featureGroup(markers)
      map!.fitBounds(group.getBounds().pad(0.1))
    } catch (error) {
      // Ignore bounds fitting errors
    }
  }
}

onUnmounted(() => {
  if (map) {
    map.remove()
    map = null
  }
})
</script>
