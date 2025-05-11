export interface Session {
  id?: number;
  thumbnail?: Thumbnail;
  sessionName?: string;
  subtitle?: string;
  description?: string;
  published?: boolean;
  date?: Date;
}

export interface Thumbnail {
  id?: number;
  filename: string;
  mimeType: string;
  data: string;
  rawData: Blob;
  cameraModel?: string;
  make?: string;
  lensModel?: string;
  exposure?: string;
  dateTime?: string;
  aperture?: number;
  iso?: number;
  fc?: number;
}
