import { AppSidebar } from "@/components/app-sidebar";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Separator } from "@/components/ui/separator";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import { SiteHeader } from "@/components/site-header";
import { ChartAreaInteractive } from "@/components/chart-area-interactive";
import { SectionCards } from "@/components/section-cards";
import { DataTable } from "@/components/data-table";
import UploadsChart from "@/components/charts/UploadsChart";
import TenantUsageChart from "@/components/charts/TenantUsageChart";
import StorageUsageChart from "@/components/charts/StorageUsageChart";
import { ChartContainer } from "@/components/charts/ChartContainer";
import FileTypeChart from "@/components/charts/FileTypeChart";
import PublicVsPrivateChart from "@/components/charts/PublicVsPrivateChart";
import SharingTrendChart from "@/components/charts/SharingTrendChart";
import SharedFilesPage from "../files/@shared/page";

export default function Page() {
  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <header className="flex h-16 shrink-0 items-center gap-2">
          <div className="flex items-center gap-2 px-4">
            <SidebarTrigger className="-ml-1" />
            <Separator
              orientation="vertical"
              className="mr-2 data-[orientation=vertical]:h-4"
            />
            <Breadcrumb>
              <BreadcrumbList>
                <BreadcrumbItem className="hidden md:block">
                  <BreadcrumbLink href="/dashboard">Home</BreadcrumbLink>
                </BreadcrumbItem>
                <BreadcrumbSeparator className="hidden md:block" />
                <BreadcrumbItem>
                  <BreadcrumbPage>Data Fetching</BreadcrumbPage>
                </BreadcrumbItem>
              </BreadcrumbList>
            </Breadcrumb>
          </div>
        </header>

        <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
          <div className="grid auto-rows-min gap-4 md:grid-cols-3">
            <div className="bg-muted/50 aspect-video rounded-xl">
              <ChartContainer>
                <StorageUsageChart />
              </ChartContainer>
            </div>
            <div className="bg-muted/50 aspect-video rounded-xl">
              <ChartContainer>
                <PublicVsPrivateChart />
              </ChartContainer>
            </div>
            <div className="bg-muted/50 aspect-video rounded-xl">
              <ChartContainer>
                <UploadsChart />
              </ChartContainer>
            </div>
          </div>
          <div className="bg-muted/50 min-h-[100vh] flex-1 rounded-xl md:min-h-min">
            <SharedFilesPage />
          </div>
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
}
