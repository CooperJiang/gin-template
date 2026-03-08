import { MockNav } from './components/MockNav'
import { MockButtons } from './components/MockButtons'
import { MockCards } from './components/MockCards'
import { MockForm } from './components/MockForm'
import { MockTable } from './components/MockTable'
import { MockAlerts } from './components/MockAlerts'
import { MockBadges } from './components/MockBadges'

export function MockPage() {
  return (
    <div data-theme-var="bg" className="min-h-full bg-gray-50">
      <MockNav />
      <div className="max-w-4xl mx-auto p-6 space-y-6">
        <MockCards />
        <div className="grid grid-cols-2 gap-6">
          <MockButtons />
          <MockBadges />
        </div>
        <MockForm />
        <MockTable />
        <MockAlerts />
      </div>
    </div>
  )
}
