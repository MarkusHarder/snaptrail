export interface Session {
  id?: string;
  thumbnail?: Thumbnail;
  sessionName?: string;
  subtitle?: string;
  description?: string;
  published?: boolean;
  date?: Date;
}

export interface Thumbnail {
  id?: string;
  filename: string;
  mimeType: string;
  data: string;
  rawData: Blob;
  imageSrc: string;
  cameraModel?: string;
  make?: string;
  lensModel?: string;
  exposure?: string;
  dateTime?: string;
  aperture?: number;
  iso?: number;
  fc?: number;
}
