#!/bin/sh

test_description="go-i18n test"

. ./lib/sharness.sh

# Create manifest repositories
manifest_url="file://${REPO_TEST_REPOSITORIES}/hello/manifests"

PODIR="$SHARNESS_TEST_DIRECTORY/../po"

test_expect_success "po-i18n without po files" '
	env LC_ALL=C go-i18n $PODIR/bad/dir >out 2>&1 &&
	grep -v "^# DEBUG" out >actual &&
	cat >expect <<-EOF &&
		############################################################
		Show messages for lang: 
		############################################################

		Hello, world.
		Welcome: guest.
		added 1 path
		added 2 paths
		added 3 paths

		############################################################
		Show messages for lang: American English
		############################################################

		Hello, world.
		Welcome: guest.
		added 1 path
		added 2 paths
		added 3 paths

		############################################################
		Show messages for lang: français
		############################################################

		Hello, world.
		Welcome: guest.
		added 1 path
		added 2 paths
		added 3 paths

		############################################################
		Show messages for lang: 中文
		############################################################

		Hello, world.
		Welcome: guest.
		added 1 path
		added 2 paths
		added 3 paths

		############################################################
		Show messages for lang: 繁體中文
		############################################################

		Hello, world.
		Welcome: guest.
		added 1 path
		added 2 paths
		added 3 paths

	EOF
	test_cmp expect actual
'

test_expect_success "po-i18n with po files" '
	env LC_ALL=C go-i18n $PODIR/build/locale >out 2>&1 &&
	grep -v "^# DEBUG" out >actual &&
	cat >expect <<-EOF &&
		############################################################
		Show messages for lang: 
		############################################################

		Hello, world.
		Welcome: guest.
		added 1 path
		added 2 paths
		added 3 paths

		############################################################
		Show messages for lang: American English
		############################################################

		Hello, world.
		Welcome: guest.
		added 1 path
		added 2 paths
		added 3 paths

		############################################################
		Afficher les messages pour la langue: français
		############################################################

		Bonjour le monde.
		Bienvenue: invité.
		ajouté 1 chemin
		ajouté 2 chemins
		ajouté 3 chemins

		############################################################
		以 中文 显示信息如下
		############################################################

		世界，您好
		欢迎：客人。
		已添加 1 个路径
		已添加 2 个路径
		已添加 3 个路径

		############################################################
		以 繁體中文 顯示訊息如下
		############################################################

		星球，您好嘛？
		歡迎：顧客。
		已添加 1 個路徑
		已添加 2 個路徑
		已添加 3 個路徑

	EOF
	test_cmp expect actual
'

test_expect_success "po-i18n say hello to user (with po files)" '
	env LC_ALL=C go-i18n $PODIR/build/locale "Jiang Xin" >out 2>&1 &&
	grep -v "^# DEBUG" out >actual &&
	cat >expect <<-EOF &&
		############################################################
		Show messages for lang: 
		############################################################

		Hello, world.
		Welcome: Jiang Xin.
		added 1 path
		added 2 paths
		added 3 paths

		############################################################
		Show messages for lang: American English
		############################################################

		Hello, world.
		Welcome: Jiang Xin.
		added 1 path
		added 2 paths
		added 3 paths

		############################################################
		Afficher les messages pour la langue: français
		############################################################

		Bonjour le monde.
		Bienvenue: Jiang Xin.
		ajouté 1 chemin
		ajouté 2 chemins
		ajouté 3 chemins

		############################################################
		以 中文 显示信息如下
		############################################################

		世界，您好
		欢迎：Jiang Xin。
		已添加 1 个路径
		已添加 2 个路径
		已添加 3 个路径

		############################################################
		以 繁體中文 顯示訊息如下
		############################################################

		星球，您好嘛？
		歡迎：Jiang Xin。
		已添加 1 個路徑
		已添加 2 個路徑
		已添加 3 個路徑

	EOF
	test_cmp expect actual
'

test_done
