import { config } from '@tamagui/config/v3'
import { Stack } from 'expo-router'
import { createTamagui, TamaguiProvider } from 'tamagui'

const tamaguiConfig = createTamagui(config)

// TypeScript types across all Tamagui APIs
type Conf = typeof tamaguiConfig

declare module 'tamagui' {
  interface TamaguiCustomConfig extends Conf {}
}

const RootLayout = () => {
  return (
    <TamaguiProvider config={tamaguiConfig}>
      <Stack screenOptions={{ title: 'Home' }} />
    </TamaguiProvider>
  )
}

export default RootLayout
