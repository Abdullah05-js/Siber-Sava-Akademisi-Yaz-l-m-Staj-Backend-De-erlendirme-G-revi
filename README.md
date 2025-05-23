# Yazılım Stajı Backend Değerlendirme Görevi

Bu proje, staj başvurusu kapsamında backend değerlendirmesi için geliştirilmiştir. Proje Go dili kullanılarak geliştirilmiş ve HTTPS destekli olarak canlı sunucuya dağıtılmıştır (hostingdunyam.com.tr).

## 🔗 API URL

Tüm endpointlere aşağıdaki domain üzerinden ulaşabilirsiniz:

https://thodex.live

### Yetkilendirme

| POST | `/api/auth/login` | API için giriş |

### TO-DO List İşlemleri

JWT token gereklidir.

| GET    | `/api/todoList/`                          | Tüm yapılacaklar listelerini getirir |

| POST   | `/api/todoList/create`                    | Yeni bir TO-DO listesi oluşturur     |

| DELETE | `/api/todoList/:todoListId`               | Belirli bir TO-DO listesini siler    |

| POST   | `/api/todoList/:todoListId/steps/create`  | Listeye yeni bir adım (step) ekler   |

| DELETE | `/api/todoList/steps/:stepId`             | Belirli bir adımı siler              |

| PATCH | /api/todoList/:todoListId/steps/:stepId | Adımı günceller |

Not: Bu endpoint, aşağıdaki gibi bir veri bekler:
{
    "content": "hello1",
    "iscomplete": false
}

Her iki alanı da ya da birini gönderebilirsiniz. Sadece gönderdiğiniz alandaki değer değişecektir.



- Domain, **Cloudflare CDN** ve güvenlik hizmeti ile desteklenmektedir.
- HTTPS trafiği, **Let's Encrypt** sertifikaları ile **Nginx** üzerinden güvenli şekilde yönlendirilmektedir.

## 👥 Kullanıcı Hesapları

Aşağıda, test amaçlı oluşturulmuş iki kullanıcı hesabı bulunmaktadır:

### 1. Normal Kullanıcı

Username: Abdullah
Password: 1,@3A
IsAdmin: false

### 2. Yönetici Kullanıcı

Username: Admin
Password: A2K,2@S
IsAdmin: true

Bu kullanıcılar ile API'yi test edebilir, login olup token alabilir ve yetkilere göre ilgili işlemleri gerçekleştirebilirsiniz.

## ⚙️ Teknik Bilgiler

- Backend dili: **Go**
- Framework **Fiber**
- HTTP sunucusu: **Nginx**
- SSL/TLS: **Let's Encrypt** kullanılarak otomatik sertifika sağlanmaktadır.
- Domain yönlendirmesi: **Cloudflare** üzerinden yapılmaktadır.
- Proje VPS üzerinde barındırılmakta ve canlı olarak çalışmaktadır.

## 📌 Notlar

- Kod yapısı, okunabilirliği ve sürdürülebilirliği ön planda tutacak şekilde sade ve anlaşılır olarak yazılmıştır.
- Gereken yerlerde açıklayıcı yorum satırları kullanılmıştır, ancak genel yaklaşım: "**İyi kod, kendini açıklar**".
- Model verilerine erişim yöntemleri optimize edilmiştir; diziler yerine map yapısı kullanılarak zaman karmaşıklığı O(n)'den O(1)'e düşürülmüştür. Bu sayede, veriye erişim işlemleri yaklaşık **n kat daha hızlı** hale getirilmiştir.
