export interface SiloSpec {
  id: string
  logoUrl?: string
  name: string
  schema: string
}

export interface SiloDefinition {
  id?: string,
  name?: string,
  siloSpecification?: SiloSpec,
  siloConfig?: object
}

type TentativeValue = 'CREATED' | 'DELETED' | null;

export interface Property {
  id?: string
  name?: string
  tentative?: TentativeValue
  categories?: Category[]
  dataSource?: DataSource
}

export interface DataSource {
  id?: string
  name?: string
  group?: string
  properties?: Property[]
  tentative?: TentativeValue
}

export interface Category {
  id?: string
  name?: string
  tentative?: TentativeValue
}

export interface Job {
  id?: string,
  jobType?: string,
  status?: 'RUNNING' | 'QUEUED' | 'FAILED' | 'COMPLETED'
  createdAt?: string
}

// Discovery models
export interface NewDataSourceDiscoveryData {
  name: string
  group?: string
  properties?: NewPropertyDiscoveryData[]
}

export interface NewPropertyDiscoveryData {
  name: string
  dataSourceId?: string
  dataSource?: DataSource | null
  categories?: NewCategoryDiscoveryData[]
}

export interface NewCategoryDiscoveryData {
  propertyId?: string
  categoryId: string
  category: Category
  property?: Property | null
}

export interface PropertyMissingDiscoveryData {
  id?: string
  property?: Property | null
}

export interface DataSourceMissingDiscoveryData {
  id?: string
  dataSource?: DataSource | null
}

type CoreDataDiscovery = {
  id?: string
  createdAt?: string
  status?: 'OPEN' | 'ACCEPTED' | 'REJECTED'
};

type NewDataSourceDiscovery = CoreDataDiscovery & {
  type?: 'DATA_SOURCE_FOUND'
  data?: NewDataSourceDiscoveryData
};

type NewPropertyDiscovery = CoreDataDiscovery & {
  type?: 'PROPERTY_FOUND'
  data?: NewPropertyDiscoveryData
};

type NewCategoryDiscovery = CoreDataDiscovery & {
  type?: 'CATEGORY_FOUND'
  data?: NewCategoryDiscoveryData
};

type MissingDataSourceDiscovery = CoreDataDiscovery & {
  data?: DataSourceMissingDiscoveryData,
  type?: 'DATA_SOURCE_MISSING'
};

type MissingPropertyDiscovery = CoreDataDiscovery & {
  data?: PropertyMissingDiscoveryData,
  type?: 'PROPERTY_MISSING'
};

export type DataDiscovery = (
  NewDataSourceDiscovery |
  NewPropertyDiscovery |
  NewCategoryDiscovery |
  MissingDataSourceDiscovery |
  MissingPropertyDiscovery
);
