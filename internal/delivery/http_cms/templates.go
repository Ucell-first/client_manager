package httpcms

import (
	"fmt"
	"net/http"

	"github.com/Ucell/client_manager/storage/repo"
)

// Oddiy HTML template funksiyalar
func (h *Handler) renderSimpleListPage(w http.ResponseWriter, users []*repo.User) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := `
<!DOCTYPE html>
<html lang="uz">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Foydalanuvchilar - Client Manager</title>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.9.4/css/bulma.min.css" rel="stylesheet">
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css" rel="stylesheet">
    <style>
        .hero { background: linear-gradient(45deg, #667eea 0%, #764ba2 100%); }
        .card { box-shadow: 0 4px 8px rgba(0,0,0,0.1); }
    </style>
</head>
<body>
    <section class="hero is-primary is-small">
        <div class="hero-body">
            <div class="container">
                <h1 class="title">
                    <i class="fas fa-users"></i> Client Manager
                </h1>
                <p class="subtitle">Foydalanuvchilar boshqaruv tizimi</p>
            </div>
        </div>
    </section>

    <section class="section">
        <div class="container">
            <div class="card">
                <header class="card-header">
                    <p class="card-header-title">
                        <i class="fas fa-list"></i> &nbsp; Foydalanuvchilar ro'yxati
                    </p>
                    <div class="card-header-icon">
                        <a href="/user/new" class="button is-primary">
                            <span class="icon"><i class="fas fa-plus"></i></span>
                            <span>Yangi qo'shish</span>
                        </a>
                    </div>
                </header>
                <div class="card-content">
`

	if len(users) > 0 {
		html += `
                    <div class="table-container">
                        <table class="table is-fullwidth is-striped is-hoverable">
                            <thead>
                                <tr>
                                    <th>ID</th>
                                    <th>MSISDN</th>
                                    <th>Ism</th>
                                    <th>Holati</th>
                                    <th>Amallar</th>
                                </tr>
                            </thead>
                            <tbody>
`
		for _, user := range users {
			activeTag := `<span class="tag is-success">Faol</span>`
			if !user.IsActive {
				activeTag = `<span class="tag is-danger">Nofaol</span>`
			}

			html += fmt.Sprintf(`
                                <tr>
                                    <td>%s</td>
                                    <td><strong>%s</strong></td>
                                    <td>%s</td>
                                    <td>%s</td>
                                    <td>
                                        <div class="buttons are-small">
                                            <a href="/user/view?id=%s" class="button is-info is-light">
                                                <i class="fas fa-eye"></i> Ko'rish
                                            </a>
                                            <a href="/user/edit?id=%s" class="button is-warning is-light">
                                                <i class="fas fa-edit"></i> Tahrirlash
                                            </a>
                                            <a href="/user/delete?id=%s"
                                               class="button is-danger is-light"
                                               onclick="return confirm('Haqiqatan ham o\\'chirmoqchimisiz?')">
                                                <i class="fas fa-trash"></i> O'chirish
                                            </a>
                                        </div>
                                    </td>
                                </tr>
			`, user.UserID, user.MSISDN, user.Name, activeTag, user.UserID, user.UserID, user.UserID)
		}
		html += `
                            </tbody>
                        </table>
                    </div>
`
	} else {
		html += `
                    <div class="notification is-info">
                        <i class="fas fa-info-circle"></i> Hozircha foydalanuvchilar yo'q.
                        <a href="/user/new">Birinchi foydalanuvchini qo'shing</a>
                    </div>
`
	}

	html += `
                </div>
            </div>
        </div>
    </section>
</body>
</html>
`

	fmt.Fprint(w, html)
}

func (h *Handler) renderSimpleNewPage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := `
<!DOCTYPE html>
<html lang="uz">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Yangi foydalanuvchi - Client Manager</title>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.9.4/css/bulma.min.css" rel="stylesheet">
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css" rel="stylesheet">
    <style>
        .hero { background: linear-gradient(45deg, #667eea 0%, #764ba2 100%); }
        .card { box-shadow: 0 4px 8px rgba(0,0,0,0.1); }
    </style>
</head>
<body>
    <section class="hero is-primary is-small">
        <div class="hero-body">
            <div class="container">
                <h1 class="title">
                    <i class="fas fa-users"></i> Client Manager
                </h1>
                <p class="subtitle">Foydalanuvchilar boshqaruv tizimi</p>
            </div>
        </div>
    </section>

    <section class="section">
        <div class="container">
            <div class="card">
                <header class="card-header">
                    <p class="card-header-title">
                        <i class="fas fa-user-plus"></i> &nbsp; Yangi foydalanuvchi qo'shish
                    </p>
                    <div class="card-header-icon">
                        <a href="/users" class="button is-light">
                            <i class="fas fa-arrow-left"></i> Orqaga
                        </a>
                    </div>
                </header>
                <div class="card-content">
                    <form method="POST" action="/user/create">
                        <div class="field">
                            <label class="label">MSISDN *</label>
                            <div class="control has-icons-left">
                                <input class="input" type="tel" name="msisdn" placeholder="+998901234567" required>
                                <span class="icon is-small is-left">
                                    <i class="fas fa-phone"></i>
                                </span>
                            </div>
                            <p class="help">Telefon raqamni to'liq kiriting (+998 bilan)</p>
                        </div>

                        <div class="field">
                            <label class="label">Ism *</label>
                            <div class="control has-icons-left">
                                <input class="input" type="text" name="name" placeholder="Foydalanuvchi ismi" required>
                                <span class="icon is-small is-left">
                                    <i class="fas fa-user"></i>
                                </span>
                            </div>
                        </div>

                        <div class="field">
                            <div class="control">
                                <label class="checkbox">
                                    <input type="checkbox" name="is_active" value="true" checked>
                                    Faol foydalanuvchi
                                </label>
                            </div>
                            <p class="help">Belgilangan bo'lsa foydalanuvchi faol bo'ladi</p>
                        </div>

                        <div class="field is-grouped">
                            <div class="control">
                                <button type="submit" class="button is-primary">
                                    <i class="fas fa-save"></i> Saqlash
                                </button>
                            </div>
                            <div class="control">
                                <a href="/users" class="button is-light">
                                    <i class="fas fa-times"></i> Bekor qilish
                                </a>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </section>
</body>
</html>
`

	fmt.Fprint(w, html)
}

func (h *Handler) renderSimpleViewPage(w http.ResponseWriter, user *repo.User) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	activeStatus := `<span class="tag is-success is-large">
		<i class="fas fa-check-circle"></i> &nbsp; Faol
	</span>`
	if !user.IsActive {
		activeStatus = `<span class="tag is-danger is-large">
			<i class="fas fa-times-circle"></i> &nbsp; Nofaol
		</span>`
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="uz">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Foydalanuvchi ma'lumotlari - Client Manager</title>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.9.4/css/bulma.min.css" rel="stylesheet">
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css" rel="stylesheet">
    <style>
        .hero { background: linear-gradient(45deg, #667eea 0%, #764ba2 100%); }
        .card { box-shadow: 0 4px 8px rgba(0,0,0,0.1); }
    </style>
</head>
<body>
    <section class="hero is-primary is-small">
        <div class="hero-body">
            <div class="container">
                <h1 class="title">
                    <i class="fas fa-users"></i> Client Manager
                </h1>
                <p class="subtitle">Foydalanuvchilar boshqaruv tizimi</p>
            </div>
        </div>
    </section>

    <section class="section">
        <div class="container">
            <div class="card">
                <header class="card-header">
                    <p class="card-header-title">
                        <i class="fas fa-user"></i> &nbsp; Foydalanuvchi ma'lumotlari
                    </p>
                    <div class="card-header-icon">
                        <a href="/users" class="button is-light">
                            <i class="fas fa-arrow-left"></i> Orqaga
                        </a>
                    </div>
                </header>
                <div class="card-content">
                    <div class="columns">
                        <div class="column is-8">
                            <table class="table is-fullwidth">
                                <tbody>
                                    <tr>
                                        <td><strong>ID:</strong></td>
                                        <td>%s</td>
                                    </tr>
                                    <tr>
                                        <td><strong>MSISDN:</strong></td>
                                        <td><code>%s</code></td>
                                    </tr>
                                    <tr>
                                        <td><strong>Ism:</strong></td>
                                        <td>%s</td>
                                    </tr>
                                    <tr>
                                        <td><strong>Holati:</strong></td>
                                        <td>%s</td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>

                    <div class="buttons">
                        <a href="/user/edit?id=%s" class="button is-warning">
                            <i class="fas fa-edit"></i> Tahrirlash
                        </a>
                        <a href="/user/delete?id=%s"
                           class="button is-danger"
                           onclick="return confirm('Haqiqatan ham o\\'chirmoqchimisiz?')">
                            <i class="fas fa-trash"></i> O'chirish
                        </a>
                        <a href="/users" class="button is-light">
                            <i class="fas fa-list"></i> Ro'yxatga qaytish
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </section>
</body>
</html>
`, user.UserID, user.MSISDN, user.Name, activeStatus, user.UserID, user.UserID)

	fmt.Fprint(w, html)
}

func (h *Handler) renderSimpleEditPage(w http.ResponseWriter, user *repo.User) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	checked := ""
	if user.IsActive {
		checked = "checked"
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="uz">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Foydalanuvchini tahrirlash - Client Manager</title>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.9.4/css/bulma.min.css" rel="stylesheet">
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css" rel="stylesheet">
    <style>
        .hero { background: linear-gradient(45deg, #667eea 0%, #764ba2 100%); }
        .card { box-shadow: 0 4px 8px rgba(0,0,0,0.1); }
    </style>
</head>
<body>
    <section class="hero is-primary is-small">
        <div class="hero-body">
            <div class="container">
                <h1 class="title">
                    <i class="fas fa-users"></i> Client Manager
                </h1>
                <p class="subtitle">Foydalanuvchilar boshqaruv tizimi</p>
            </div>
        </div>
    </section>

    <section class="section">
        <div class="container">
            <div class="card">
                <header class="card-header">
                    <p class="card-header-title">
                        <i class="fas fa-user-edit"></i> &nbsp; Foydalanuvchini tahrirlash
                    </p>
                    <div class="card-header-icon">
                        <a href="/users" class="button is-light">
                            <i class="fas fa-arrow-left"></i> Orqaga
                        </a>
                    </div>
                </header>
                <div class="card-content">
                    <form method="POST" action="/user/update">
                        <input type="hidden" name="user_id" value="%s">

                        <div class="field">
                            <label class="label">User ID</label>
                            <div class="control">
                                <input class="input" type="text" value="%s" disabled>
                            </div>
                            <p class="help">ID o'zgartirilmaydi</p>
                        </div>

                        <div class="field">
                            <label class="label">MSISDN *</label>
                            <div class="control has-icons-left">
                                <input class="input" type="tel" name="msisdn" value="%s" required>
                                <span class="icon is-small is-left">
                                    <i class="fas fa-phone"></i>
                                </span>
                            </div>
                        </div>

                        <div class="field">
                            <label class="label">Ism *</label>
                            <div class="control has-icons-left">
                                <input class="input" type="text" name="name" value="%s" required>
                                <span class="icon is-small is-left">
                                    <i class="fas fa-user"></i>
                                </span>
                            </div>
                        </div>

                        <div class="field">
                            <div class="control">
                                <label class="checkbox">
                                    <input type="checkbox" name="is_active" value="true" %s>
                                    Faol foydalanuvchi
                                </label>
                            </div>
                        </div>

                        <div class="field is-grouped">
                            <div class="control">
                                <button type="submit" class="button is-warning">
                                    <i class="fas fa-save"></i> Yangilash
                                </button>
                            </div>
                            <div class="control">
                                <a href="/user/view?id=%s" class="button is-info">
                                    <i class="fas fa-eye"></i> Ko'rish
                                </a>
                            </div>
                            <div class="control">
                                <a href="/users" class="button is-light">
                                    <i class="fas fa-times"></i> Bekor qilish
                                </a>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </section>
</body>
</html>
`, user.UserID, user.UserID, user.MSISDN, user.Name, checked, user.UserID)

	fmt.Fprint(w, html)
}
