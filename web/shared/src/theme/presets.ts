import type { ThemeDefinition } from './types'
import { defaultLightColors, defaultDarkColors, defaultTokens } from './defaults'

function deepClone<T>(obj: T): T {
  return JSON.parse(JSON.stringify(obj))
}

const greenLight = deepClone(defaultLightColors)
greenLight.primary = {
  50: '240 253 244', 100: '220 252 231', 200: '187 247 208', 300: '134 239 172',
  400: '74 222 128', 500: '34 197 94', 600: '22 163 74', 700: '21 128 61',
  800: '22 101 52', 900: '20 83 45', 950: '5 46 22',
}
greenLight.ring = '34 197 94'
greenLight.link = '22 163 74'
greenLight.linkHover = '21 128 61'

const greenDark = deepClone(defaultDarkColors)
greenDark.primary = {
  50: '5 46 22', 100: '20 83 45', 200: '22 101 52', 300: '21 128 61',
  400: '22 163 74', 500: '34 197 94', 600: '74 222 128', 700: '134 239 172',
  800: '187 247 208', 900: '220 252 231', 950: '240 253 244',
}
greenDark.ring = '74 222 128'
greenDark.link = '74 222 128'
greenDark.linkHover = '134 239 172'

const purpleLight = deepClone(defaultLightColors)
purpleLight.primary = {
  50: '250 245 255', 100: '243 232 255', 200: '233 213 255', 300: '216 180 254',
  400: '192 132 252', 500: '168 85 247', 600: '147 51 234', 700: '126 34 206',
  800: '107 33 168', 900: '88 28 135', 950: '59 7 100',
}
purpleLight.ring = '168 85 247'
purpleLight.link = '147 51 234'
purpleLight.linkHover = '126 34 206'

const purpleDark = deepClone(defaultDarkColors)
purpleDark.primary = {
  50: '59 7 100', 100: '88 28 135', 200: '107 33 168', 300: '126 34 206',
  400: '147 51 234', 500: '168 85 247', 600: '192 132 252', 700: '216 180 254',
  800: '233 213 255', 900: '243 232 255', 950: '250 245 255',
}
purpleDark.ring = '192 132 252'
purpleDark.link = '192 132 252'
purpleDark.linkHover = '216 180 254'

const orangeLight = deepClone(defaultLightColors)
orangeLight.primary = {
  50: '255 247 237', 100: '255 237 213', 200: '254 215 170', 300: '253 186 116',
  400: '251 146 60', 500: '249 115 22', 600: '234 88 12', 700: '194 65 12',
  800: '154 52 18', 900: '124 45 18', 950: '67 20 7',
}
orangeLight.ring = '249 115 22'
orangeLight.link = '234 88 12'
orangeLight.linkHover = '194 65 12'

const orangeDark = deepClone(defaultDarkColors)
orangeDark.primary = {
  50: '67 20 7', 100: '124 45 18', 200: '154 52 18', 300: '194 65 12',
  400: '234 88 12', 500: '249 115 22', 600: '251 146 60', 700: '253 186 116',
  800: '254 215 170', 900: '255 237 213', 950: '255 247 237',
}
orangeDark.ring = '251 146 60'
orangeDark.link = '251 146 60'
orangeDark.linkHover = '253 186 116'

const roseLight = deepClone(defaultLightColors)
roseLight.primary = {
  50: '255 241 242', 100: '255 228 230', 200: '254 205 211', 300: '253 164 175',
  400: '251 113 133', 500: '244 63 94', 600: '225 29 72', 700: '190 18 60',
  800: '159 18 57', 900: '136 19 55', 950: '76 5 25',
}
roseLight.ring = '244 63 94'
roseLight.link = '225 29 72'
roseLight.linkHover = '190 18 60'

const roseDark = deepClone(defaultDarkColors)
roseDark.primary = {
  50: '76 5 25', 100: '136 19 55', 200: '159 18 57', 300: '190 18 60',
  400: '225 29 72', 500: '244 63 94', 600: '251 113 133', 700: '253 164 175',
  800: '254 205 211', 900: '255 228 230', 950: '255 241 242',
}
roseDark.ring = '251 113 133'
roseDark.link = '251 113 133'
roseDark.linkHover = '253 164 175'

export const presetThemes: ThemeDefinition[] = [
  { name: 'Default Blue', light: defaultLightColors, dark: defaultDarkColors, tokens: defaultTokens },
  { name: 'Green', light: greenLight, dark: greenDark, tokens: defaultTokens },
  { name: 'Purple', light: purpleLight, dark: purpleDark, tokens: defaultTokens },
  { name: 'Orange', light: orangeLight, dark: orangeDark, tokens: defaultTokens },
  { name: 'Rose', light: roseLight, dark: roseDark, tokens: defaultTokens },
]
