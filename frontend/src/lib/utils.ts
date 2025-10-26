import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function capitalizeFirstWords(str: string): string {
  return str.replace(/\b\w/g, (char) => char.toUpperCase());
}
